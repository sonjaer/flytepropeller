package nodes

import (
	"context"
	"fmt"
	"strconv"

	"github.com/lyft/flyteidl/clients/go/events"
	"github.com/lyft/flytestdlib/storage"
	"k8s.io/apimachinery/pkg/types"

	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/io"
	"github.com/lyft/flyteplugins/go/tasks/pluginmachinery/ioutils"

	"github.com/lyft/flytepropeller/pkg/apis/flyteworkflow/v1alpha1"
	"github.com/lyft/flytepropeller/pkg/controller/nodes/handler"
	"github.com/lyft/flytepropeller/pkg/utils"
)

const NodeIDLabel = "node-id"
const TaskNameLabel = "task-name"
const NodeInterruptibleLabel = "interruptible"

type execMetadata struct {
	v1alpha1.WorkflowMeta
	interrutptible bool
	nodeLabels     map[string]string
}

func (e execMetadata) GetK8sServiceAccount() string {
	return e.WorkflowMeta.GetServiceAccountName()
}

func (e execMetadata) GetOwnerID() types.NamespacedName {
	return types.NamespacedName{Name: e.GetName(), Namespace: e.GetNamespace()}
}

func (e execMetadata) IsInterruptible() bool {
	return e.interrutptible
}

func (e execMetadata) GetLabels() map[string]string {
	return e.nodeLabels
}

type execContext struct {
	store               *storage.DataStore
	tr                  handler.TaskReader
	md                  handler.NodeExecutionMetadata
	er                  events.TaskEventRecorder
	inputs              io.InputReader
	node                v1alpha1.ExecutableNode
	nodeStatus          v1alpha1.ExecutableNodeStatus
	maxDatasetSizeBytes int64
	nsm                 *nodeStateManager
	enqueueOwner        func() error
	w                   v1alpha1.ExecutableWorkflow
	rawOutputPrefix     storage.DataReference
	shardSelector       ioutils.ShardSelector
}

func (e execContext) OutputShardSelector() ioutils.ShardSelector {
	return e.shardSelector
}

func (e execContext) RawOutputPrefix() storage.DataReference {
	return e.rawOutputPrefix
}

func (e execContext) EnqueueOwnerFunc() func() error {
	return e.enqueueOwner
}

func (e execContext) Workflow() v1alpha1.ExecutableWorkflow {
	return e.w
}

func (e execContext) TaskReader() handler.TaskReader {
	return e.tr
}

func (e execContext) NodeStateReader() handler.NodeStateReader {
	return e.nsm
}

func (e execContext) NodeStateWriter() handler.NodeStateWriter {
	return e.nsm
}

func (e execContext) DataStore() *storage.DataStore {
	return e.store
}

func (e execContext) InputReader() io.InputReader {
	return e.inputs
}

func (e execContext) EventsRecorder() events.TaskEventRecorder {
	return e.er
}

func (e execContext) NodeID() v1alpha1.NodeID {
	return e.node.GetID()
}

func (e execContext) Node() v1alpha1.ExecutableNode {
	return e.node
}

func (e execContext) CurrentAttempt() uint32 {
	return e.nodeStatus.GetAttempts()
}

func (e execContext) NodeStatus() v1alpha1.ExecutableNodeStatus {
	return e.nodeStatus
}

func (e execContext) NodeExecutionMetadata() handler.NodeExecutionMetadata {
	return e.md
}

func (e execContext) MaxDatasetSizeBytes() int64 {
	return e.maxDatasetSizeBytes
}

func newNodeExecContext(_ context.Context, store *storage.DataStore, w v1alpha1.ExecutableWorkflow, node v1alpha1.ExecutableNode, nodeStatus v1alpha1.ExecutableNodeStatus, inputs io.InputReader, interruptible bool, maxDatasetSize int64, er events.TaskEventRecorder, tr handler.TaskReader, nsm *nodeStateManager, enqueueOwner func() error, rawOutputPrefix storage.DataReference, outputShardSelector ioutils.ShardSelector) *execContext {
	md := execMetadata{WorkflowMeta: w, interrutptible: interruptible}

	// Copy the wf labels before adding node specific labels.
	nodeLabels := make(map[string]string)
	for k, v := range w.GetLabels() {
		nodeLabels[k] = v
	}
	nodeLabels[NodeIDLabel] = utils.SanitizeLabelValue(node.GetID())
	if tr != nil && tr.GetTaskID() != nil {
		nodeLabels[TaskNameLabel] = utils.SanitizeLabelValue(tr.GetTaskID().Name)
	}
	nodeLabels[NodeInterruptibleLabel] = strconv.FormatBool(interruptible)
	md.nodeLabels = nodeLabels

	return &execContext{
		md:                  md,
		store:               store,
		node:                node,
		nodeStatus:          nodeStatus,
		inputs:              inputs,
		er:                  er,
		maxDatasetSizeBytes: maxDatasetSize,
		tr:                  tr,
		nsm:                 nsm,
		enqueueOwner:        enqueueOwner,
		w:                   w,
		rawOutputPrefix:     rawOutputPrefix,
		shardSelector:       outputShardSelector,
	}
}

func (c *nodeExecutor) newNodeExecContextDefault(ctx context.Context, w v1alpha1.ExecutableWorkflow, n v1alpha1.ExecutableNode, s v1alpha1.ExecutableNodeStatus) (*execContext, error) {
	var tr handler.TaskReader
	if n.GetKind() == v1alpha1.NodeKindTask {
		if n.GetTaskID() == nil {
			return nil, fmt.Errorf("bad state, no task-id defined for node [%s]", n.GetID())
		}
		tk, err := w.GetTask(*n.GetTaskID())
		if err != nil {
			return nil, err
		}
		tr = &taskReader{TaskTemplate: tk.CoreTask()}
	}

	workflowEnqueuer := func() error {
		c.enqueueWorkflow(w.GetID())
		return nil
	}

	interrutible := w.IsInterruptible()
	if n.IsInterruptible() != nil {
		interrutible = *n.IsInterruptible()
	}

	// a node is not considered interruptible if the system failures have exceeded the configured threshold
	if interrutible && s.GetSystemFailures() >= c.interruptibleFailureThreshold {
		interrutible = false
		c.metrics.InterruptedThresholdHit.Inc(ctx)
	}

	return newNodeExecContext(ctx, c.store, w, n, s,
		ioutils.NewCachedInputReader(
			ctx,
			ioutils.NewRemoteFileInputReader(
				ctx,
				c.store,
				ioutils.NewInputFilePaths(
					ctx,
					c.store,
					s.GetDataDir(),
				),
			),
		),
		interrutible,
		c.maxDatasetSizeBytes,
		&taskEventRecorder{TaskEventRecorder: c.taskRecorder},
		tr,
		newNodeStateManager(ctx, s),
		workflowEnqueuer,
		// Eventually we want to replace this with per workflow sandboxes
		// https://github.com/lyft/flyte/issues/211
		c.defaultDataSandbox,
		c.shardSelector,
	), nil
}
