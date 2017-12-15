// +build !ignore_autogenerated

/*
Copyright 2017 The KubeVirt Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package api

import (
	time "time"

	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"

	api_v1 "kubevirt.io/kubevirt/pkg/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Address) DeepCopyInto(out *Address) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Address.
func (in *Address) DeepCopy() *Address {
	if in == nil {
		return nil
	}
	out := new(Address)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alias) DeepCopyInto(out *Alias) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alias.
func (in *Alias) DeepCopy() *Alias {
	if in == nil {
		return nil
	}
	out := new(Alias)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BIOS) DeepCopyInto(out *BIOS) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BIOS.
func (in *BIOS) DeepCopy() *BIOS {
	if in == nil {
		return nil
	}
	out := new(BIOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Ballooning) DeepCopyInto(out *Ballooning) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Ballooning.
func (in *Ballooning) DeepCopy() *Ballooning {
	if in == nil {
		return nil
	}
	out := new(Ballooning)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BandWidth) DeepCopyInto(out *BandWidth) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BandWidth.
func (in *BandWidth) DeepCopy() *BandWidth {
	if in == nil {
		return nil
	}
	out := new(BandWidth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Boot) DeepCopyInto(out *Boot) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Boot.
func (in *Boot) DeepCopy() *Boot {
	if in == nil {
		return nil
	}
	out := new(Boot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootMenu) DeepCopyInto(out *BootMenu) {
	*out = *in
	if in.Timeout != nil {
		in, out := &in.Timeout, &out.Timeout
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootMenu.
func (in *BootMenu) DeepCopy() *BootMenu {
	if in == nil {
		return nil
	}
	out := new(BootMenu)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BootOrder) DeepCopyInto(out *BootOrder) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BootOrder.
func (in *BootOrder) DeepCopy() *BootOrder {
	if in == nil {
		return nil
	}
	out := new(BootOrder)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Channel) DeepCopyInto(out *Channel) {
	*out = *in
	out.Source = in.Source
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		if *in == nil {
			*out = nil
		} else {
			*out = new(ChannelTarget)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Channel.
func (in *Channel) DeepCopy() *Channel {
	if in == nil {
		return nil
	}
	out := new(Channel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelSource) DeepCopyInto(out *ChannelSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelSource.
func (in *ChannelSource) DeepCopy() *ChannelSource {
	if in == nil {
		return nil
	}
	out := new(ChannelSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChannelTarget) DeepCopyInto(out *ChannelTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChannelTarget.
func (in *ChannelTarget) DeepCopy() *ChannelTarget {
	if in == nil {
		return nil
	}
	out := new(ChannelTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Clock) DeepCopyInto(out *Clock) {
	*out = *in
	if in.Timer != nil {
		in, out := &in.Timer, &out.Timer
		*out = make([]Timer, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Clock.
func (in *Clock) DeepCopy() *Clock {
	if in == nil {
		return nil
	}
	out := new(Clock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Commandline) DeepCopyInto(out *Commandline) {
	*out = *in
	if in.QEMUEnv != nil {
		in, out := &in.QEMUEnv, &out.QEMUEnv
		*out = make([]Env, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Commandline.
func (in *Commandline) DeepCopy() *Commandline {
	if in == nil {
		return nil
	}
	out := new(Commandline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Console) DeepCopyInto(out *Console) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		if *in == nil {
			*out = nil
		} else {
			*out = new(ConsoleTarget)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Console.
func (in *Console) DeepCopy() *Console {
	if in == nil {
		return nil
	}
	out := new(Console)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConsoleTarget) DeepCopyInto(out *ConsoleTarget) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		if *in == nil {
			*out = nil
		} else {
			*out = new(string)
			**out = **in
		}
	}
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConsoleTarget.
func (in *ConsoleTarget) DeepCopy() *ConsoleTarget {
	if in == nil {
		return nil
	}
	out := new(ConsoleTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Context) DeepCopyInto(out *Context) {
	*out = *in
	if in.Secrets != nil {
		in, out := &in.Secrets, &out.Secrets
		*out = make(map[string]*v1.Secret, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(v1.Secret)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	if in.VirtualMachine != nil {
		in, out := &in.VirtualMachine, &out.VirtualMachine
		if *in == nil {
			*out = nil
		} else {
			*out = new(api_v1.VirtualMachine)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Context.
func (in *Context) DeepCopy() *Context {
	if in == nil {
		return nil
	}
	out := new(Context)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Devices) DeepCopyInto(out *Devices) {
	*out = *in
	if in.Interfaces != nil {
		in, out := &in.Interfaces, &out.Interfaces
		*out = make([]Interface, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Channels != nil {
		in, out := &in.Channels, &out.Channels
		*out = make([]Channel, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Video != nil {
		in, out := &in.Video, &out.Video
		*out = make([]Video, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Graphics != nil {
		in, out := &in.Graphics, &out.Graphics
		*out = make([]Graphics, len(*in))
		copy(*out, *in)
	}
	if in.Ballooning != nil {
		in, out := &in.Ballooning, &out.Ballooning
		if *in == nil {
			*out = nil
		} else {
			*out = new(Ballooning)
			**out = **in
		}
	}
	if in.Disks != nil {
		in, out := &in.Disks, &out.Disks
		*out = make([]Disk, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Serials != nil {
		in, out := &in.Serials, &out.Serials
		*out = make([]Serial, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Consoles != nil {
		in, out := &in.Consoles, &out.Consoles
		*out = make([]Console, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Watchdog != nil {
		in, out := &in.Watchdog, &out.Watchdog
		if *in == nil {
			*out = nil
		} else {
			*out = new(Watchdog)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Devices.
func (in *Devices) DeepCopy() *Devices {
	if in == nil {
		return nil
	}
	out := new(Devices)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Disk) DeepCopyInto(out *Disk) {
	*out = *in
	in.Source.DeepCopyInto(&out.Source)
	out.Target = in.Target
	if in.Driver != nil {
		in, out := &in.Driver, &out.Driver
		if *in == nil {
			*out = nil
		} else {
			*out = new(DiskDriver)
			**out = **in
		}
	}
	if in.ReadOnly != nil {
		in, out := &in.ReadOnly, &out.ReadOnly
		if *in == nil {
			*out = nil
		} else {
			*out = new(ReadOnly)
			**out = **in
		}
	}
	if in.Auth != nil {
		in, out := &in.Auth, &out.Auth
		if *in == nil {
			*out = nil
		} else {
			*out = new(DiskAuth)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		if *in == nil {
			*out = nil
		} else {
			*out = new(Alias)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Disk.
func (in *Disk) DeepCopy() *Disk {
	if in == nil {
		return nil
	}
	out := new(Disk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskAuth) DeepCopyInto(out *DiskAuth) {
	*out = *in
	if in.Secret != nil {
		in, out := &in.Secret, &out.Secret
		if *in == nil {
			*out = nil
		} else {
			*out = new(DiskSecret)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskAuth.
func (in *DiskAuth) DeepCopy() *DiskAuth {
	if in == nil {
		return nil
	}
	out := new(DiskAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskDriver) DeepCopyInto(out *DiskDriver) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskDriver.
func (in *DiskDriver) DeepCopy() *DiskDriver {
	if in == nil {
		return nil
	}
	out := new(DiskDriver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSecret) DeepCopyInto(out *DiskSecret) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSecret.
func (in *DiskSecret) DeepCopy() *DiskSecret {
	if in == nil {
		return nil
	}
	out := new(DiskSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSource) DeepCopyInto(out *DiskSource) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		if *in == nil {
			*out = nil
		} else {
			*out = new(DiskSourceHost)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSource.
func (in *DiskSource) DeepCopy() *DiskSource {
	if in == nil {
		return nil
	}
	out := new(DiskSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskSourceHost) DeepCopyInto(out *DiskSourceHost) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskSourceHost.
func (in *DiskSourceHost) DeepCopy() *DiskSourceHost {
	if in == nil {
		return nil
	}
	out := new(DiskSourceHost)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DiskTarget) DeepCopyInto(out *DiskTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DiskTarget.
func (in *DiskTarget) DeepCopy() *DiskTarget {
	if in == nil {
		return nil
	}
	out := new(DiskTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Domain) DeepCopyInto(out *Domain) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Domain.
func (in *Domain) DeepCopy() *Domain {
	if in == nil {
		return nil
	}
	out := new(Domain)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Domain) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainList) DeepCopyInto(out *DomainList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Domain, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainList.
func (in *DomainList) DeepCopy() *DomainList {
	if in == nil {
		return nil
	}
	out := new(DomainList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DomainList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainSpec) DeepCopyInto(out *DomainSpec) {
	*out = *in
	out.XMLName = in.XMLName
	out.Memory = in.Memory
	in.OS.DeepCopyInto(&out.OS)
	if in.SysInfo != nil {
		in, out := &in.SysInfo, &out.SysInfo
		if *in == nil {
			*out = nil
		} else {
			*out = new(SysInfo)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Devices.DeepCopyInto(&out.Devices)
	if in.Clock != nil {
		in, out := &in.Clock, &out.Clock
		if *in == nil {
			*out = nil
		} else {
			*out = new(Clock)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Resource != nil {
		in, out := &in.Resource, &out.Resource
		if *in == nil {
			*out = nil
		} else {
			*out = new(Resource)
			**out = **in
		}
	}
	if in.QEMUCmd != nil {
		in, out := &in.QEMUCmd, &out.QEMUCmd
		if *in == nil {
			*out = nil
		} else {
			*out = new(Commandline)
			(*in).DeepCopyInto(*out)
		}
	}
	in.Metadata.DeepCopyInto(&out.Metadata)
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		if *in == nil {
			*out = nil
		} else {
			*out = new(Features)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainSpec.
func (in *DomainSpec) DeepCopy() *DomainSpec {
	if in == nil {
		return nil
	}
	out := new(DomainSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DomainStatus) DeepCopyInto(out *DomainStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DomainStatus.
func (in *DomainStatus) DeepCopy() *DomainStatus {
	if in == nil {
		return nil
	}
	out := new(DomainStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Entry) DeepCopyInto(out *Entry) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Entry.
func (in *Entry) DeepCopy() *Entry {
	if in == nil {
		return nil
	}
	out := new(Entry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Env) DeepCopyInto(out *Env) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Env.
func (in *Env) DeepCopy() *Env {
	if in == nil {
		return nil
	}
	out := new(Env)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureEnabled) DeepCopyInto(out *FeatureEnabled) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureEnabled.
func (in *FeatureEnabled) DeepCopy() *FeatureEnabled {
	if in == nil {
		return nil
	}
	out := new(FeatureEnabled)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureHyperv) DeepCopyInto(out *FeatureHyperv) {
	*out = *in
	if in.Relaxed != nil {
		in, out := &in.Relaxed, &out.Relaxed
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			**out = **in
		}
	}
	if in.VAPIC != nil {
		in, out := &in.VAPIC, &out.VAPIC
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			**out = **in
		}
	}
	if in.Spinlocks != nil {
		in, out := &in.Spinlocks, &out.Spinlocks
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureSpinlocks)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.VPIndex != nil {
		in, out := &in.VPIndex, &out.VPIndex
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			**out = **in
		}
	}
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			**out = **in
		}
	}
	if in.SyNIC != nil {
		in, out := &in.SyNIC, &out.SyNIC
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			**out = **in
		}
	}
	if in.SyNICTimer != nil {
		in, out := &in.SyNICTimer, &out.SyNICTimer
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			**out = **in
		}
	}
	if in.Reset != nil {
		in, out := &in.Reset, &out.Reset
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureState)
			**out = **in
		}
	}
	if in.VendorID != nil {
		in, out := &in.VendorID, &out.VendorID
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureVendorID)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureHyperv.
func (in *FeatureHyperv) DeepCopy() *FeatureHyperv {
	if in == nil {
		return nil
	}
	out := new(FeatureHyperv)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureSpinlocks) DeepCopyInto(out *FeatureSpinlocks) {
	*out = *in
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureSpinlocks.
func (in *FeatureSpinlocks) DeepCopy() *FeatureSpinlocks {
	if in == nil {
		return nil
	}
	out := new(FeatureSpinlocks)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureState) DeepCopyInto(out *FeatureState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureState.
func (in *FeatureState) DeepCopy() *FeatureState {
	if in == nil {
		return nil
	}
	out := new(FeatureState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureVendorID) DeepCopyInto(out *FeatureVendorID) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureVendorID.
func (in *FeatureVendorID) DeepCopy() *FeatureVendorID {
	if in == nil {
		return nil
	}
	out := new(FeatureVendorID)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Features) DeepCopyInto(out *Features) {
	*out = *in
	if in.ACPI != nil {
		in, out := &in.ACPI, &out.ACPI
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureEnabled)
			**out = **in
		}
	}
	if in.APIC != nil {
		in, out := &in.APIC, &out.APIC
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureEnabled)
			**out = **in
		}
	}
	if in.Hyperv != nil {
		in, out := &in.Hyperv, &out.Hyperv
		if *in == nil {
			*out = nil
		} else {
			*out = new(FeatureHyperv)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Features.
func (in *Features) DeepCopy() *Features {
	if in == nil {
		return nil
	}
	out := new(Features)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FilterRef) DeepCopyInto(out *FilterRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FilterRef.
func (in *FilterRef) DeepCopy() *FilterRef {
	if in == nil {
		return nil
	}
	out := new(FilterRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GracePeriodMetadata) DeepCopyInto(out *GracePeriodMetadata) {
	*out = *in
	if in.DeletionTimestamp != nil {
		in, out := &in.DeletionTimestamp, &out.DeletionTimestamp
		if *in == nil {
			*out = nil
		} else {
			*out = new(time.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GracePeriodMetadata.
func (in *GracePeriodMetadata) DeepCopy() *GracePeriodMetadata {
	if in == nil {
		return nil
	}
	out := new(GracePeriodMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Graphics) DeepCopyInto(out *Graphics) {
	*out = *in
	out.Listen = in.Listen
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Graphics.
func (in *Graphics) DeepCopy() *Graphics {
	if in == nil {
		return nil
	}
	out := new(Graphics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Interface) DeepCopyInto(out *Interface) {
	*out = *in
	if in.Address != nil {
		in, out := &in.Address, &out.Address
		if *in == nil {
			*out = nil
		} else {
			*out = new(Address)
			**out = **in
		}
	}
	out.Source = in.Source
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		if *in == nil {
			*out = nil
		} else {
			*out = new(InterfaceTarget)
			**out = **in
		}
	}
	if in.Model != nil {
		in, out := &in.Model, &out.Model
		if *in == nil {
			*out = nil
		} else {
			*out = new(Model)
			**out = **in
		}
	}
	if in.MAC != nil {
		in, out := &in.MAC, &out.MAC
		if *in == nil {
			*out = nil
		} else {
			*out = new(MAC)
			**out = **in
		}
	}
	if in.BandWidth != nil {
		in, out := &in.BandWidth, &out.BandWidth
		if *in == nil {
			*out = nil
		} else {
			*out = new(BandWidth)
			**out = **in
		}
	}
	if in.BootOrder != nil {
		in, out := &in.BootOrder, &out.BootOrder
		if *in == nil {
			*out = nil
		} else {
			*out = new(BootOrder)
			**out = **in
		}
	}
	if in.LinkState != nil {
		in, out := &in.LinkState, &out.LinkState
		if *in == nil {
			*out = nil
		} else {
			*out = new(LinkState)
			**out = **in
		}
	}
	if in.FilterRef != nil {
		in, out := &in.FilterRef, &out.FilterRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(FilterRef)
			**out = **in
		}
	}
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		if *in == nil {
			*out = nil
		} else {
			*out = new(Alias)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Interface.
func (in *Interface) DeepCopy() *Interface {
	if in == nil {
		return nil
	}
	out := new(Interface)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceSource) DeepCopyInto(out *InterfaceSource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceSource.
func (in *InterfaceSource) DeepCopy() *InterfaceSource {
	if in == nil {
		return nil
	}
	out := new(InterfaceSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InterfaceTarget) DeepCopyInto(out *InterfaceTarget) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InterfaceTarget.
func (in *InterfaceTarget) DeepCopy() *InterfaceTarget {
	if in == nil {
		return nil
	}
	out := new(InterfaceTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinkState) DeepCopyInto(out *LinkState) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinkState.
func (in *LinkState) DeepCopy() *LinkState {
	if in == nil {
		return nil
	}
	out := new(LinkState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Listen) DeepCopyInto(out *Listen) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Listen.
func (in *Listen) DeepCopy() *Listen {
	if in == nil {
		return nil
	}
	out := new(Listen)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Loader) DeepCopyInto(out *Loader) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Loader.
func (in *Loader) DeepCopy() *Loader {
	if in == nil {
		return nil
	}
	out := new(Loader)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MAC) DeepCopyInto(out *MAC) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MAC.
func (in *MAC) DeepCopy() *MAC {
	if in == nil {
		return nil
	}
	out := new(MAC)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Memory) DeepCopyInto(out *Memory) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Memory.
func (in *Memory) DeepCopy() *Memory {
	if in == nil {
		return nil
	}
	out := new(Memory)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	in.GracePeriod.DeepCopyInto(&out.GracePeriod)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Model) DeepCopyInto(out *Model) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Model.
func (in *Model) DeepCopy() *Model {
	if in == nil {
		return nil
	}
	out := new(Model)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NVRam) DeepCopyInto(out *NVRam) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NVRam.
func (in *NVRam) DeepCopy() *NVRam {
	if in == nil {
		return nil
	}
	out := new(NVRam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OS) DeepCopyInto(out *OS) {
	*out = *in
	out.Type = in.Type
	if in.SMBios != nil {
		in, out := &in.SMBios, &out.SMBios
		if *in == nil {
			*out = nil
		} else {
			*out = new(SMBios)
			**out = **in
		}
	}
	if in.BootOrder != nil {
		in, out := &in.BootOrder, &out.BootOrder
		*out = make([]Boot, len(*in))
		copy(*out, *in)
	}
	if in.BootMenu != nil {
		in, out := &in.BootMenu, &out.BootMenu
		if *in == nil {
			*out = nil
		} else {
			*out = new(BootMenu)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.BIOS != nil {
		in, out := &in.BIOS, &out.BIOS
		if *in == nil {
			*out = nil
		} else {
			*out = new(BIOS)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OS.
func (in *OS) DeepCopy() *OS {
	if in == nil {
		return nil
	}
	out := new(OS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OSType) DeepCopyInto(out *OSType) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OSType.
func (in *OSType) DeepCopy() *OSType {
	if in == nil {
		return nil
	}
	out := new(OSType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RandomGenerator) DeepCopyInto(out *RandomGenerator) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RandomGenerator.
func (in *RandomGenerator) DeepCopy() *RandomGenerator {
	if in == nil {
		return nil
	}
	out := new(RandomGenerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReadOnly) DeepCopyInto(out *ReadOnly) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReadOnly.
func (in *ReadOnly) DeepCopy() *ReadOnly {
	if in == nil {
		return nil
	}
	out := new(ReadOnly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Resource) DeepCopyInto(out *Resource) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Resource.
func (in *Resource) DeepCopy() *Resource {
	if in == nil {
		return nil
	}
	out := new(Resource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SMBios) DeepCopyInto(out *SMBios) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SMBios.
func (in *SMBios) DeepCopy() *SMBios {
	if in == nil {
		return nil
	}
	out := new(SMBios)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretSpec) DeepCopyInto(out *SecretSpec) {
	*out = *in
	out.XMLName = in.XMLName
	out.Usage = in.Usage
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretSpec.
func (in *SecretSpec) DeepCopy() *SecretSpec {
	if in == nil {
		return nil
	}
	out := new(SecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretUsage) DeepCopyInto(out *SecretUsage) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretUsage.
func (in *SecretUsage) DeepCopy() *SecretUsage {
	if in == nil {
		return nil
	}
	out := new(SecretUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Serial) DeepCopyInto(out *Serial) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		if *in == nil {
			*out = nil
		} else {
			*out = new(SerialTarget)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Serial.
func (in *Serial) DeepCopy() *Serial {
	if in == nil {
		return nil
	}
	out := new(Serial)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SerialTarget) DeepCopyInto(out *SerialTarget) {
	*out = *in
	if in.Port != nil {
		in, out := &in.Port, &out.Port
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SerialTarget.
func (in *SerialTarget) DeepCopy() *SerialTarget {
	if in == nil {
		return nil
	}
	out := new(SerialTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SysInfo) DeepCopyInto(out *SysInfo) {
	*out = *in
	if in.System != nil {
		in, out := &in.System, &out.System
		*out = make([]Entry, len(*in))
		copy(*out, *in)
	}
	if in.BIOS != nil {
		in, out := &in.BIOS, &out.BIOS
		*out = make([]Entry, len(*in))
		copy(*out, *in)
	}
	if in.BaseBoard != nil {
		in, out := &in.BaseBoard, &out.BaseBoard
		*out = make([]Entry, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SysInfo.
func (in *SysInfo) DeepCopy() *SysInfo {
	if in == nil {
		return nil
	}
	out := new(SysInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timer) DeepCopyInto(out *Timer) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timer.
func (in *Timer) DeepCopy() *Timer {
	if in == nil {
		return nil
	}
	out := new(Timer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Video) DeepCopyInto(out *Video) {
	*out = *in
	in.Model.DeepCopyInto(&out.Model)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Video.
func (in *Video) DeepCopy() *Video {
	if in == nil {
		return nil
	}
	out := new(Video)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VideoModel) DeepCopyInto(out *VideoModel) {
	*out = *in
	if in.Heads != nil {
		in, out := &in.Heads, &out.Heads
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint)
			**out = **in
		}
	}
	if in.Ram != nil {
		in, out := &in.Ram, &out.Ram
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint)
			**out = **in
		}
	}
	if in.VRam != nil {
		in, out := &in.VRam, &out.VRam
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint)
			**out = **in
		}
	}
	if in.VGAMem != nil {
		in, out := &in.VGAMem, &out.VGAMem
		if *in == nil {
			*out = nil
		} else {
			*out = new(uint)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VideoModel.
func (in *VideoModel) DeepCopy() *VideoModel {
	if in == nil {
		return nil
	}
	out := new(VideoModel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Watchdog) DeepCopyInto(out *Watchdog) {
	*out = *in
	if in.Alias != nil {
		in, out := &in.Alias, &out.Alias
		if *in == nil {
			*out = nil
		} else {
			*out = new(Alias)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Watchdog.
func (in *Watchdog) DeepCopy() *Watchdog {
	if in == nil {
		return nil
	}
	out := new(Watchdog)
	in.DeepCopyInto(out)
	return out
}
