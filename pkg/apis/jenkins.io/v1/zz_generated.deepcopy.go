// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommitSummary) DeepCopyInto(out *CommitSummary) {
	*out = *in
	if in.Author != nil {
		in, out := &in.Author, &out.Author
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Committer != nil {
		in, out := &in.Committer, &out.Committer
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommitSummary.
func (in *CommitSummary) DeepCopy() *CommitSummary {
	if in == nil {
		return nil
	}
	out := new(CommitSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CoreActivityStep) DeepCopyInto(out *CoreActivityStep) {
	*out = *in
	if in.StartedTimestamp != nil {
		in, out := &in.StartedTimestamp, &out.StartedTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.CompletedTimestamp != nil {
		in, out := &in.CompletedTimestamp, &out.CompletedTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CoreActivityStep.
func (in *CoreActivityStep) DeepCopy() *CoreActivityStep {
	if in == nil {
		return nil
	}
	out := new(CoreActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Environment) DeepCopyInto(out *Environment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Environment.
func (in *Environment) DeepCopy() *Environment {
	if in == nil {
		return nil
	}
	out := new(Environment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Environment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentList) DeepCopyInto(out *EnvironmentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Environment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentList.
func (in *EnvironmentList) DeepCopy() *EnvironmentList {
	if in == nil {
		return nil
	}
	out := new(EnvironmentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EnvironmentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentRepository) DeepCopyInto(out *EnvironmentRepository) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentRepository.
func (in *EnvironmentRepository) DeepCopy() *EnvironmentRepository {
	if in == nil {
		return nil
	}
	out := new(EnvironmentRepository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentSpec) DeepCopyInto(out *EnvironmentSpec) {
	*out = *in
	out.Source = in.Source
	out.TeamSettings = in.TeamSettings
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentSpec.
func (in *EnvironmentSpec) DeepCopy() *EnvironmentSpec {
	if in == nil {
		return nil
	}
	out := new(EnvironmentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentStatus) DeepCopyInto(out *EnvironmentStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentStatus.
func (in *EnvironmentStatus) DeepCopy() *EnvironmentStatus {
	if in == nil {
		return nil
	}
	out := new(EnvironmentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitStatus) DeepCopyInto(out *GitStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitStatus.
func (in *GitStatus) DeepCopy() *GitStatus {
	if in == nil {
		return nil
	}
	out := new(GitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IssueSummary) DeepCopyInto(out *IssueSummary) {
	*out = *in
	if in.User != nil {
		in, out := &in.User, &out.User
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Assignees != nil {
		in, out := &in.Assignees, &out.Assignees
		*out = make([]UserDetails, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ClosedBy != nil {
		in, out := &in.ClosedBy, &out.ClosedBy
		if *in == nil {
			*out = nil
		} else {
			*out = new(UserDetails)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IssueSummary.
func (in *IssueSummary) DeepCopy() *IssueSummary {
	if in == nil {
		return nil
	}
	out := new(IssueSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivity) DeepCopyInto(out *PipelineActivity) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivity.
func (in *PipelineActivity) DeepCopy() *PipelineActivity {
	if in == nil {
		return nil
	}
	out := new(PipelineActivity)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineActivity) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivityList) DeepCopyInto(out *PipelineActivityList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PipelineActivity, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivityList.
func (in *PipelineActivityList) DeepCopy() *PipelineActivityList {
	if in == nil {
		return nil
	}
	out := new(PipelineActivityList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PipelineActivityList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivitySpec) DeepCopyInto(out *PipelineActivitySpec) {
	*out = *in
	if in.StartedTimestamp != nil {
		in, out := &in.StartedTimestamp, &out.StartedTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.CompletedTimestamp != nil {
		in, out := &in.CompletedTimestamp, &out.CompletedTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]PipelineActivityStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivitySpec.
func (in *PipelineActivitySpec) DeepCopy() *PipelineActivitySpec {
	if in == nil {
		return nil
	}
	out := new(PipelineActivitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivityStatus) DeepCopyInto(out *PipelineActivityStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivityStatus.
func (in *PipelineActivityStatus) DeepCopy() *PipelineActivityStatus {
	if in == nil {
		return nil
	}
	out := new(PipelineActivityStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PipelineActivityStep) DeepCopyInto(out *PipelineActivityStep) {
	*out = *in
	if in.Stage != nil {
		in, out := &in.Stage, &out.Stage
		if *in == nil {
			*out = nil
		} else {
			*out = new(StageActivityStep)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Promote != nil {
		in, out := &in.Promote, &out.Promote
		if *in == nil {
			*out = nil
		} else {
			*out = new(PromoteActivityStep)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PipelineActivityStep.
func (in *PipelineActivityStep) DeepCopy() *PipelineActivityStep {
	if in == nil {
		return nil
	}
	out := new(PipelineActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PreviewGitSpec) DeepCopyInto(out *PreviewGitSpec) {
	*out = *in
	out.User = in.User
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PreviewGitSpec.
func (in *PreviewGitSpec) DeepCopy() *PreviewGitSpec {
	if in == nil {
		return nil
	}
	out := new(PreviewGitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromoteActivityStep) DeepCopyInto(out *PromoteActivityStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		if *in == nil {
			*out = nil
		} else {
			*out = new(PromotePullRequestStep)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Update != nil {
		in, out := &in.Update, &out.Update
		if *in == nil {
			*out = nil
		} else {
			*out = new(PromoteUpdateStep)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromoteActivityStep.
func (in *PromoteActivityStep) DeepCopy() *PromoteActivityStep {
	if in == nil {
		return nil
	}
	out := new(PromoteActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromotePullRequestStep) DeepCopyInto(out *PromotePullRequestStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromotePullRequestStep.
func (in *PromotePullRequestStep) DeepCopy() *PromotePullRequestStep {
	if in == nil {
		return nil
	}
	out := new(PromotePullRequestStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PromoteUpdateStep) DeepCopyInto(out *PromoteUpdateStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	if in.Statuses != nil {
		in, out := &in.Statuses, &out.Statuses
		*out = make([]GitStatus, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PromoteUpdateStep.
func (in *PromoteUpdateStep) DeepCopy() *PromoteUpdateStep {
	if in == nil {
		return nil
	}
	out := new(PromoteUpdateStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Release) DeepCopyInto(out *Release) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Release.
func (in *Release) DeepCopy() *Release {
	if in == nil {
		return nil
	}
	out := new(Release)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Release) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseList) DeepCopyInto(out *ReleaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Release, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseList.
func (in *ReleaseList) DeepCopy() *ReleaseList {
	if in == nil {
		return nil
	}
	out := new(ReleaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReleaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseSpec) DeepCopyInto(out *ReleaseSpec) {
	*out = *in
	if in.Commits != nil {
		in, out := &in.Commits, &out.Commits
		*out = make([]CommitSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Issues != nil {
		in, out := &in.Issues, &out.Issues
		*out = make([]IssueSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PullRequests != nil {
		in, out := &in.PullRequests, &out.PullRequests
		*out = make([]IssueSummary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseSpec.
func (in *ReleaseSpec) DeepCopy() *ReleaseSpec {
	if in == nil {
		return nil
	}
	out := new(ReleaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReleaseStatus) DeepCopyInto(out *ReleaseStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReleaseStatus.
func (in *ReleaseStatus) DeepCopy() *ReleaseStatus {
	if in == nil {
		return nil
	}
	out := new(ReleaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StageActivityStep) DeepCopyInto(out *StageActivityStep) {
	*out = *in
	in.CoreActivityStep.DeepCopyInto(&out.CoreActivityStep)
	if in.Steps != nil {
		in, out := &in.Steps, &out.Steps
		*out = make([]CoreActivityStep, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StageActivityStep.
func (in *StageActivityStep) DeepCopy() *StageActivityStep {
	if in == nil {
		return nil
	}
	out := new(StageActivityStep)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TeamSettings) DeepCopyInto(out *TeamSettings) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TeamSettings.
func (in *TeamSettings) DeepCopy() *TeamSettings {
	if in == nil {
		return nil
	}
	out := new(TeamSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserDetails) DeepCopyInto(out *UserDetails) {
	*out = *in
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = (*in).DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserDetails.
func (in *UserDetails) DeepCopy() *UserDetails {
	if in == nil {
		return nil
	}
	out := new(UserDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserSpec) DeepCopyInto(out *UserSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserSpec.
func (in *UserSpec) DeepCopy() *UserSpec {
	if in == nil {
		return nil
	}
	out := new(UserSpec)
	in.DeepCopyInto(out)
	return out
}
