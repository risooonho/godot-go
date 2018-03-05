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

//func NewipFromPointer(ptr gdnative.Pointer) ip {
func NewIPFromPointer(ptr gdnative.Pointer) ip {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ip{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonIP() *ip {
	return &ip{}
}

/*
   IP contains support functions for the Internet Protocol (IP). TCP/IP support is in different classes (see [StreamPeerTCP] and [TCP_Server]). IP provides DNS hostname resolution support, both blocking and threaded.
*/
var IP = newSingletonIP()

/*
IP contains support functions for the Internet Protocol (IP). TCP/IP support is in different classes (see [StreamPeerTCP] and [TCP_Server]). IP provides DNS hostname resolution support, both blocking and threaded.
*/
type ip struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *ip) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("IP")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *ip) BaseClass() string {
	return "IP"
}

/*
        Removes all of a "hostname"'s cached references. If no "hostname" is given then all cached IP addresses are removed.
	Args: [{ true hostname String}], Returns: void
*/
func (o *ip) ClearCache(hostname gdnative.String) {
	o.ensureSingleton()
	//log.Println("Calling IP.ClearCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(hostname)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("IP", "clear_cache")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes a given item "id" from the queue. This should be used to free a queue after it has completed to enable more queries to happen.
	Args: [{ false id int}], Returns: void
*/
func (o *ip) EraseResolveItem(id gdnative.Int) {
	o.ensureSingleton()
	//log.Println("Calling IP.EraseResolveItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("IP", "erase_resolve_item")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns all of the user's current IPv4 and IPv6 addresses as an array.
	Args: [], Returns: Array
*/
func (o *ip) GetLocalAddresses() gdnative.Array {
	o.ensureSingleton()
	//log.Println("Calling IP.GetLocalAddresses()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("IP", "get_local_addresses")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns a queued hostname's IP address, given its queue "id". Returns an empty string on error or if resolution hasn't happened yet (see [method get_resolve_item_status]).
	Args: [{ false id int}], Returns: String
*/
func (o *ip) GetResolveItemAddress(id gdnative.Int) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling IP.GetResolveItemAddress()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("IP", "get_resolve_item_address")

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
        Returns a queued hostname's status as a RESOLVER_STATUS_* constant, given its queue "id".
	Args: [{ false id int}], Returns: enum.IP::ResolverStatus
*/

/*
        Returns a given hostname's IPv4 or IPv6 address when resolved (blocking-type method). The address type returned depends on the TYPE_* constant given as "ip_type".
	Args: [{ false host String} {3 true ip_type int}], Returns: String
*/
func (o *ip) ResolveHostname(host gdnative.String, ipType gdnative.Int) gdnative.String {
	o.ensureSingleton()
	//log.Println("Calling IP.ResolveHostname()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(host)
	ptrArguments[1] = gdnative.NewPointerFromInt(ipType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("IP", "resolve_hostname")

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
        Creates a queue item to resolve a hostname to an IPv4 or IPv6 address depending on the TYPE_* constant given as "ip_type". Returns the queue ID if successful, or RESOLVER_INVALID_ID on error.
	Args: [{ false host String} {3 true ip_type int}], Returns: int
*/
func (o *ip) ResolveHostnameQueueItem(host gdnative.String, ipType gdnative.Int) gdnative.Int {
	o.ensureSingleton()
	//log.Println("Calling IP.ResolveHostnameQueueItem()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(host)
	ptrArguments[1] = gdnative.NewPointerFromInt(ipType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("IP", "resolve_hostname_queue_item")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}
