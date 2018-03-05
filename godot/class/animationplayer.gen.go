package class

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewAnimationPlayerFromPointer(ptr gdnative.Pointer) AnimationPlayer {
func NewAnimationPlayerFromPointer(ptr gdnative.Pointer) AnimationPlayer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationPlayer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
An animation player is used for general purpose playback of [Animation] resources. It contains a dictionary of animations (referenced by name) and custom blend times between their transitions. Additionally, animations can be played and blended in different channels.
*/
type AnimationPlayer struct {
	Node
	owner gdnative.Object
}

func (o *AnimationPlayer) BaseClass() string {
	return "AnimationPlayer"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AnimationPlayer) X_AnimationChanged() {
	//log.Println("Calling AnimationPlayer.X_AnimationChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "_animation_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Object}], Returns: void
*/
func (o *AnimationPlayer) X_NodeRemoved(arg0 Object) {
	//log.Println("Calling AnimationPlayer.X_NodeRemoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "_node_removed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds [code]animation[/code] to the player accessible with the key [code]name[/code].
	Args: [{ false name String} { false animation Animation}], Returns: enum.Error
*/

/*
        Shifts position in the animation timeline. Delta is the time in seconds to shift.
	Args: [{ false delta float}], Returns: void
*/
func (o *AnimationPlayer) Advance(delta gdnative.Float) {
	//log.Println("Calling AnimationPlayer.Advance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(delta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "advance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the name of the next animation in the queue.
	Args: [{ false anim_from String}], Returns: String
*/
func (o *AnimationPlayer) AnimationGetNext(animFrom gdnative.String) gdnative.String {
	//log.Println("Calling AnimationPlayer.AnimationGetNext()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(animFrom)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "animation_get_next")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Triggers the [code]anim_to[/code] animation when the [code]anim_from[/code] animation completes.
	Args: [{ false anim_from String} { false anim_to String}], Returns: void
*/
func (o *AnimationPlayer) AnimationSetNext(animFrom gdnative.String, animTo gdnative.String) {
	//log.Println("Calling AnimationPlayer.AnimationSetNext()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(animFrom)
	ptrArguments[1] = gdnative.NewPointerFromString(animTo)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "animation_set_next")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        [code]AnimationPlayer[/code] caches animated nodes. It may not notice if a node disappears, so clear_caches forces it to update the cache again.
	Args: [], Returns: void
*/
func (o *AnimationPlayer) ClearCaches() {
	//log.Println("Calling AnimationPlayer.ClearCaches()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "clear_caches")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Clears all queued, unplayed animations.
	Args: [], Returns: void
*/
func (o *AnimationPlayer) ClearQueue() {
	//log.Println("Calling AnimationPlayer.ClearQueue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "clear_queue")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the name of [code]animation[/code] or empty string if not found.
	Args: [{ false animation Animation}], Returns: String
*/
func (o *AnimationPlayer) FindAnimation(animation Animation) gdnative.String {
	//log.Println("Calling AnimationPlayer.FindAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(animation.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "find_animation")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the [Animation] with key [code]name[/code] or [code]null[/code] if not found.
	Args: [{ false name String}], Returns: Animation
*/
func (o *AnimationPlayer) GetAnimation(name gdnative.String) Animation {
	//log.Println("Calling AnimationPlayer.GetAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_animation")

	// Call the parent method.
	// Animation
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewAnimationFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the list of stored animation names.
	Args: [], Returns: PoolStringArray
*/
func (o *AnimationPlayer) GetAnimationList() gdnative.PoolStringArray {
	//log.Println("Calling AnimationPlayer.GetAnimationList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_animation_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.AnimationPlayer::AnimationProcessMode
*/

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimationPlayer) GetAssignedAnimation() gdnative.String {
	//log.Println("Calling AnimationPlayer.GetAssignedAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_assigned_animation")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimationPlayer) GetAutoplay() gdnative.String {
	//log.Println("Calling AnimationPlayer.GetAutoplay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_autoplay")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Get the blend time (in seconds) between two animations, referenced by their names.
	Args: [{ false anim_from String} { false anim_to String}], Returns: float
*/
func (o *AnimationPlayer) GetBlendTime(animFrom gdnative.String, animTo gdnative.String) gdnative.Float {
	//log.Println("Calling AnimationPlayer.GetBlendTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(animFrom)
	ptrArguments[1] = gdnative.NewPointerFromString(animTo)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_blend_time")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *AnimationPlayer) GetCurrentAnimation() gdnative.String {
	//log.Println("Calling AnimationPlayer.GetCurrentAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_current_animation")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AnimationPlayer) GetCurrentAnimationLength() gdnative.Float {
	//log.Println("Calling AnimationPlayer.GetCurrentAnimationLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_current_animation_length")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AnimationPlayer) GetCurrentAnimationPosition() gdnative.Float {
	//log.Println("Calling AnimationPlayer.GetCurrentAnimationPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_current_animation_position")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AnimationPlayer) GetDefaultBlendTime() gdnative.Float {
	//log.Println("Calling AnimationPlayer.GetDefaultBlendTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_default_blend_time")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *AnimationPlayer) GetRoot() gdnative.NodePath {
	//log.Println("Calling AnimationPlayer.GetRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_root")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AnimationPlayer) GetSpeedScale() gdnative.Float {
	//log.Println("Calling AnimationPlayer.GetSpeedScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "get_speed_scale")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyFloat()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewFloatFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns [code]true[/code] if the [code]AnimationPlayer[/code] stores an [Animation] with key [code]name[/code].
	Args: [{ false name String}], Returns: bool
*/
func (o *AnimationPlayer) HasAnimation(name gdnative.String) gdnative.Bool {
	//log.Println("Calling AnimationPlayer.HasAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "has_animation")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AnimationPlayer) IsActive() gdnative.Bool {
	//log.Println("Calling AnimationPlayer.IsActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "is_active")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns [code]true[/code] if playing an animation.
	Args: [], Returns: bool
*/
func (o *AnimationPlayer) IsPlaying() gdnative.Bool {
	//log.Println("Calling AnimationPlayer.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "is_playing")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Play the animation with key [code]name[/code]. Custom speed and blend times can be set. If custom speed is negative (-1), 'from_end' being true can play the animation backwards.
	Args: [{ true name String} {-1 true custom_blend float} {1 true custom_speed float} {False true from_end bool}], Returns: void
*/
func (o *AnimationPlayer) Play(name gdnative.String, customBlend gdnative.Float, customSpeed gdnative.Float, fromEnd gdnative.Bool) {
	//log.Println("Calling AnimationPlayer.Play()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromFloat(customBlend)
	ptrArguments[2] = gdnative.NewPointerFromFloat(customSpeed)
	ptrArguments[3] = gdnative.NewPointerFromBool(fromEnd)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "play")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Play the animation with key [code]name[/code] in reverse.
	Args: [{ true name String} {-1 true custom_blend float}], Returns: void
*/
func (o *AnimationPlayer) PlayBackwards(name gdnative.String, customBlend gdnative.Float) {
	//log.Println("Calling AnimationPlayer.PlayBackwards()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromFloat(customBlend)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "play_backwards")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Queue an animation for playback once the current one is done.
	Args: [{ false name String}], Returns: void
*/
func (o *AnimationPlayer) Queue(name gdnative.String) {
	//log.Println("Calling AnimationPlayer.Queue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "queue")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Remove the animation with key [code]name[/code].
	Args: [{ false name String}], Returns: void
*/
func (o *AnimationPlayer) RemoveAnimation(name gdnative.String) {
	//log.Println("Calling AnimationPlayer.RemoveAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "remove_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Rename an existing animation with key [code]name[/code] to [code]newname[/code].
	Args: [{ false name String} { false newname String}], Returns: void
*/
func (o *AnimationPlayer) RenameAnimation(name gdnative.String, newname gdnative.String) {
	//log.Println("Calling AnimationPlayer.RenameAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromString(newname)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "rename_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Seek the animation to the [code]seconds[/code] point in time (in seconds). If [code]update[/code] is [code]true[/code], the animation updates too, otherwise it updates at process time.
	Args: [{ false seconds float} {False true update bool}], Returns: void
*/
func (o *AnimationPlayer) Seek(seconds gdnative.Float, update gdnative.Bool) {
	//log.Println("Calling AnimationPlayer.Seek()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromFloat(seconds)
	ptrArguments[1] = gdnative.NewPointerFromBool(update)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "seek")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false active bool}], Returns: void
*/
func (o *AnimationPlayer) SetActive(active gdnative.Bool) {
	//log.Println("Calling AnimationPlayer.SetActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(active)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_active")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *AnimationPlayer) SetAnimationProcessMode(mode gdnative.Int) {
	//log.Println("Calling AnimationPlayer.SetAnimationProcessMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_animation_process_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false anim String}], Returns: void
*/
func (o *AnimationPlayer) SetAssignedAnimation(anim gdnative.String) {
	//log.Println("Calling AnimationPlayer.SetAssignedAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_assigned_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/
func (o *AnimationPlayer) SetAutoplay(name gdnative.String) {
	//log.Println("Calling AnimationPlayer.SetAutoplay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_autoplay")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Specify a blend time (in seconds) between two animations, referenced by their names.
	Args: [{ false anim_from String} { false anim_to String} { false sec float}], Returns: void
*/
func (o *AnimationPlayer) SetBlendTime(animFrom gdnative.String, animTo gdnative.String, sec gdnative.Float) {
	//log.Println("Calling AnimationPlayer.SetBlendTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(animFrom)
	ptrArguments[1] = gdnative.NewPointerFromString(animTo)
	ptrArguments[2] = gdnative.NewPointerFromFloat(sec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_blend_time")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false anim String}], Returns: void
*/
func (o *AnimationPlayer) SetCurrentAnimation(anim gdnative.String) {
	//log.Println("Calling AnimationPlayer.SetCurrentAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(anim)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_current_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false sec float}], Returns: void
*/
func (o *AnimationPlayer) SetDefaultBlendTime(sec gdnative.Float) {
	//log.Println("Calling AnimationPlayer.SetDefaultBlendTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(sec)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_default_blend_time")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false path NodePath}], Returns: void
*/
func (o *AnimationPlayer) SetRoot(path gdnative.NodePath) {
	//log.Println("Calling AnimationPlayer.SetRoot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_root")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false speed float}], Returns: void
*/
func (o *AnimationPlayer) SetSpeedScale(speed gdnative.Float) {
	//log.Println("Calling AnimationPlayer.SetSpeedScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromFloat(speed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "set_speed_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stop the currently playing animation. If [code]reset[/code] is [code]true[/code], the anim position is reset to [code]0[/code].
	Args: [{True true reset bool}], Returns: void
*/
func (o *AnimationPlayer) Stop(reset gdnative.Bool) {
	//log.Println("Calling AnimationPlayer.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(reset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationPlayer", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}