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

//func NewImageFromPointer(ptr gdnative.Pointer) Image {
func NewImageFromPointer(ptr gdnative.Pointer) Image {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Image{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Native image datatype. Contains image data, which can be converted to a [Texture], and several functions to interact with it. The maximum width and height for an [code]Image[/code] is 16384 pixels.
*/
type Image struct {
	Resource
	owner gdnative.Object
}

func (o *Image) BaseClass() string {
	return "Image"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *Image) X_GetData() gdnative.Dictionary {
	//log.Println("Calling Image.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "_get_data")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false data Dictionary}], Returns: void
*/
func (o *Image) X_SetData(data gdnative.Dictionary) {
	//log.Println("Calling Image.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Alpha-blends [code]src_rect[/code] from [code]src[/code] image to this image at coordinates [code]dest[/code].
	Args: [{ false src Image} { false src_rect Rect2} { false dst Vector2}], Returns: void
*/
func (o *Image) BlendRect(src Image, srcRect gdnative.Rect2, dst gdnative.Vector2) {
	//log.Println("Calling Image.BlendRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(src.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromRect2(srcRect)
	ptrArguments[2] = gdnative.NewPointerFromVector2(dst)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "blend_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Alpha-blends [code]src_rect[/code] from [code]src[/code] image to this image using [code]mask[/code] image at coordinates [code]dst[/code]. Alpha channels are required for both [code]src[/code] and [code]mask[/code]. [code]dst[/code] pixels and [code]src[/code] pixels will blend if the corresponding mask pixel's alpha value is not 0. [code]src[/code] image and [code]mask[/code] image [b]must[/b] have the same size (width and height) but they can have different formats.
	Args: [{ false src Image} { false mask Image} { false src_rect Rect2} { false dst Vector2}], Returns: void
*/
func (o *Image) BlendRectMask(src Image, mask Image, srcRect gdnative.Rect2, dst gdnative.Vector2) {
	//log.Println("Calling Image.BlendRectMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(src.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromObject(mask.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromRect2(srcRect)
	ptrArguments[3] = gdnative.NewPointerFromVector2(dst)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "blend_rect_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Copies [code]src_rect[/code] from [code]src[/code] image to this image at coordinates [code]dst[/code].
	Args: [{ false src Image} { false src_rect Rect2} { false dst Vector2}], Returns: void
*/
func (o *Image) BlitRect(src Image, srcRect gdnative.Rect2, dst gdnative.Vector2) {
	//log.Println("Calling Image.BlitRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(src.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromRect2(srcRect)
	ptrArguments[2] = gdnative.NewPointerFromVector2(dst)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "blit_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Blits [code]src_rect[/code] area from [code]src[/code] image to this image at the coordinates given by [code]dst[/code]. [code]src[/code] pixel is copied onto [code]dst[/code] if the corresponding [code]mask[/code] pixel's alpha value is not 0. [code]src[/code] image and [code]mask[/code] image [b]must[/b] have the same size (width and height) but they can have different formats.
	Args: [{ false src Image} { false mask Image} { false src_rect Rect2} { false dst Vector2}], Returns: void
*/
func (o *Image) BlitRectMask(src Image, mask Image, srcRect gdnative.Rect2, dst gdnative.Vector2) {
	//log.Println("Calling Image.BlitRectMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(src.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromObject(mask.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromRect2(srcRect)
	ptrArguments[3] = gdnative.NewPointerFromVector2(dst)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "blit_rect_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes the image's mipmaps.
	Args: [], Returns: void
*/
func (o *Image) ClearMipmaps() {
	//log.Println("Calling Image.ClearMipmaps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "clear_mipmaps")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Compresses the image to use less memory. Can not directly access pixel data while the image is compressed. Returns error if the chosen compression mode is not available. See [code]COMPRESS_*[/code] constants.
	Args: [{ false mode int} { false source int} { false lossy_quality float}], Returns: enum.Error
*/

/*
        Converts the image's format. See [code]FORMAT_*[/code] constants.
	Args: [{ false format int}], Returns: void
*/
func (o *Image) Convert(format gdnative.Int) {
	//log.Println("Calling Image.Convert()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(format)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "convert")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Copies [code]src[/code] image to this image.
	Args: [{ false src Image}], Returns: void
*/
func (o *Image) CopyFrom(src Image) {
	//log.Println("Calling Image.CopyFrom()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(src.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "copy_from")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates an empty image of given size and format. See [code]FORMAT_*[/code] constants. If [code]use_mipmaps[/code] is true then generate mipmaps for this image. See the [code]generate_mipmaps[/code] method.
	Args: [{ false width int} { false height int} { false use_mipmaps bool} { false format int}], Returns: void
*/
func (o *Image) Create(width gdnative.Int, height gdnative.Int, useMipmaps gdnative.Bool, format gdnative.Int) {
	//log.Println("Calling Image.Create()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)
	ptrArguments[1] = gdnative.NewPointerFromInt(height)
	ptrArguments[2] = gdnative.NewPointerFromBool(useMipmaps)
	ptrArguments[3] = gdnative.NewPointerFromInt(format)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "create")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Creates a new image of given size and format. See [code]FORMAT_*[/code] constants. Fills the image with the given raw data. If [code]use_mipmaps[/code] is true then generate mipmaps for this image. See the [code]generate_mipmaps[/code] method.
	Args: [{ false width int} { false height int} { false use_mipmaps bool} { false format int} { false data PoolByteArray}], Returns: void
*/
func (o *Image) CreateFromData(width gdnative.Int, height gdnative.Int, useMipmaps gdnative.Bool, format gdnative.Int, data gdnative.PoolByteArray) {
	//log.Println("Calling Image.CreateFromData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)
	ptrArguments[1] = gdnative.NewPointerFromInt(height)
	ptrArguments[2] = gdnative.NewPointerFromBool(useMipmaps)
	ptrArguments[3] = gdnative.NewPointerFromInt(format)
	ptrArguments[4] = gdnative.NewPointerFromPoolByteArray(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "create_from_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Crops the image to the given [code]width[/code] and [code]height[/code]. If the specified size is larger than the current size, the extra area is filled with black pixels.
	Args: [{ false width int} { false height int}], Returns: void
*/
func (o *Image) Crop(width gdnative.Int, height gdnative.Int) {
	//log.Println("Calling Image.Crop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)
	ptrArguments[1] = gdnative.NewPointerFromInt(height)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "crop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Decompresses the image if it is compressed. Returns an error if decompress function is not available.
	Args: [], Returns: enum.Error
*/

/*
        Returns ALPHA_BLEND if the image has data for alpha values. Returns ALPHA_BIT if all the alpha values are below a certain threshold or the maximum value. Returns ALPHA_NONE if no data for alpha values is found.
	Args: [], Returns: enum.Image::AlphaMode
*/

/*
        Stretches the image and enlarges it by a factor of 2. No interpolation is done.
	Args: [], Returns: void
*/
func (o *Image) ExpandX2Hq2X() {
	//log.Println("Calling Image.ExpandX2Hq2X()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "expand_x2_hq2x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Fills the image with a given [Color].
	Args: [{ false color Color}], Returns: void
*/
func (o *Image) Fill(color gdnative.Color) {
	//log.Println("Calling Image.Fill()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "fill")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Blends low-alpha pixels with nearby pixels.
	Args: [], Returns: void
*/
func (o *Image) FixAlphaEdges() {
	//log.Println("Calling Image.FixAlphaEdges()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "fix_alpha_edges")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Flips the image horizontally.
	Args: [], Returns: void
*/
func (o *Image) FlipX() {
	//log.Println("Calling Image.FlipX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "flip_x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Flips the image vertically.
	Args: [], Returns: void
*/
func (o *Image) FlipY() {
	//log.Println("Calling Image.FlipY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "flip_y")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Generates mipmaps for the image. Mipmaps are pre-calculated and lower resolution copies of the image. Mipmaps are automatically used if the image needs to be scaled down when rendered. This improves image quality and the performance of the rendering. Returns an error if the image is compressed, in a custom format or if the image's width/height is 0.
	Args: [], Returns: enum.Error
*/

/*
        Returns the image's raw data.
	Args: [], Returns: PoolByteArray
*/
func (o *Image) GetData() gdnative.PoolByteArray {
	//log.Println("Calling Image.GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "get_data")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the image’s format. See [code]FORMAT_*[/code] constants.
	Args: [], Returns: enum.Image::Format
*/

/*
        Returns the image's height.
	Args: [], Returns: int
*/
func (o *Image) GetHeight() gdnative.Int {
	//log.Println("Calling Image.GetHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "get_height")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the offset where the image's mipmap with index [code]mipmap[/code] is stored in the [code]data[/code] dictionary.
	Args: [{ false mipmap int}], Returns: int
*/
func (o *Image) GetMipmapOffset(mipmap gdnative.Int) gdnative.Int {
	//log.Println("Calling Image.GetMipmapOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mipmap)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "get_mipmap_offset")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the color of the pixel at [code](x, y)[/code] if the image is locked. If the image is unlocked it always returns a [Color] with the value [code](0, 0, 0, 1.0)[/code].
	Args: [{ false x int} { false y int}], Returns: Color
*/
func (o *Image) GetPixel(x gdnative.Int, y gdnative.Int) gdnative.Color {
	//log.Println("Calling Image.GetPixel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(x)
	ptrArguments[1] = gdnative.NewPointerFromInt(y)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "get_pixel")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns a new image that is a copy of the image's area specified with [code]rect[/code].
	Args: [{ false rect Rect2}], Returns: Image
*/
func (o *Image) GetRect(rect gdnative.Rect2) Image {
	//log.Println("Calling Image.GetRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRect2(rect)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "get_rect")

	// Call the parent method.
	// Image
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewImageFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the image's size (width and height).
	Args: [], Returns: Vector2
*/
func (o *Image) GetSize() gdnative.Vector2 {
	//log.Println("Calling Image.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "get_size")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns a [Rect2] enclosing the visible portion of the image.
	Args: [], Returns: Rect2
*/
func (o *Image) GetUsedRect() gdnative.Rect2 {
	//log.Println("Calling Image.GetUsedRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "get_used_rect")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns the image's width.
	Args: [], Returns: int
*/
func (o *Image) GetWidth() gdnative.Int {
	//log.Println("Calling Image.GetWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "get_width")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Returns [code]true[/code] if the image has generated mipmaps.
	Args: [], Returns: bool
*/
func (o *Image) HasMipmaps() gdnative.Bool {
	//log.Println("Calling Image.HasMipmaps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "has_mipmaps")

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
        Returns [code]true[/code] if the image is compressed.
	Args: [], Returns: bool
*/
func (o *Image) IsCompressed() gdnative.Bool {
	//log.Println("Calling Image.IsCompressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "is_compressed")

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
        Returns [code]true[/code] if the image has no data.
	Args: [], Returns: bool
*/
func (o *Image) IsEmpty() gdnative.Bool {
	//log.Println("Calling Image.IsEmpty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "is_empty")

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
        Returns [code]true[/code] if all the image's pixels have an alpha value of 0. Returns [code]false[/code] if any pixel has an alpha value higher than 0.
	Args: [], Returns: bool
*/
func (o *Image) IsInvisible() gdnative.Bool {
	//log.Println("Calling Image.IsInvisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "is_invisible")

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
        Loads an image from file [code]path[/code].
	Args: [{ false path String}], Returns: enum.Error
*/

/*

	Args: [{ false buffer PoolByteArray}], Returns: enum.Error
*/

/*

	Args: [{ false buffer PoolByteArray}], Returns: enum.Error
*/

/*
        Locks the data for writing access.
	Args: [], Returns: void
*/
func (o *Image) Lock() {
	//log.Println("Calling Image.Lock()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "lock")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Converts the image's data to represent coordinates on a 3D plane. This is used when the image represents a normalmap. A normalmap can add lots of detail to a 3D surface without increasing the polygon count.
	Args: [], Returns: void
*/
func (o *Image) NormalmapToXy() {
	//log.Println("Calling Image.NormalmapToXy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "normalmap_to_xy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Multiplies color values with alpha values. Resulting color values for a pixel are [code](color * alpha)/256[/code].
	Args: [], Returns: void
*/
func (o *Image) PremultiplyAlpha() {
	//log.Println("Calling Image.PremultiplyAlpha()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "premultiply_alpha")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Resizes the image to the given [code]width[/code] and [code]height[/code]. New pixels are calculated using [code]interpolation[/code]. See [code]interpolation[/code] constants.
	Args: [{ false width int} { false height int} {1 true interpolation int}], Returns: void
*/
func (o *Image) Resize(width gdnative.Int, height gdnative.Int, interpolation gdnative.Int) {
	//log.Println("Calling Image.Resize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)
	ptrArguments[1] = gdnative.NewPointerFromInt(height)
	ptrArguments[2] = gdnative.NewPointerFromInt(interpolation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "resize")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Resizes the image to the nearest power of 2 for the width and height. If [code]square[/code] is [code]true[/code] then set width and height to be the same.
	Args: [{False true square bool}], Returns: void
*/
func (o *Image) ResizeToPo2(square gdnative.Bool) {
	//log.Println("Calling Image.ResizeToPo2()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(square)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "resize_to_po2")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Saves the image as a PNG file to [code]path[/code].
	Args: [{ false path String}], Returns: enum.Error
*/

/*
        Sets the [Color] of the pixel at [code](x, y)[/code] if the image is locked. Example: [codeblock] var img = Image.new() img.create(img_width, img_height, false, Image.FORMAT_RGBA8) img.lock() img.set_pixel(x, y, color) # Works img.unlock() img.set_pixel(x, y, color) # Does not have an effect [/codeblock]
	Args: [{ false x int} { false y int} { false color Color}], Returns: void
*/
func (o *Image) SetPixel(x gdnative.Int, y gdnative.Int, color gdnative.Color) {
	//log.Println("Calling Image.SetPixel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(x)
	ptrArguments[1] = gdnative.NewPointerFromInt(y)
	ptrArguments[2] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "set_pixel")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Shrinks the image by a factor of 2.
	Args: [], Returns: void
*/
func (o *Image) ShrinkX2() {
	//log.Println("Calling Image.ShrinkX2()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "shrink_x2")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Converts the raw data from the sRGB colorspace to a linear scale.
	Args: [], Returns: void
*/
func (o *Image) SrgbToLinear() {
	//log.Println("Calling Image.SrgbToLinear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "srgb_to_linear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Unlocks the data and prevents changes.
	Args: [], Returns: void
*/
func (o *Image) Unlock() {
	//log.Println("Calling Image.Unlock()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Image", "unlock")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}
