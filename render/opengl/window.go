package opengl

import (
	"github.com/fcvarela/gosg/core"
	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
	"github.com/golang/glog"
)

// PostMakeWindow implements the core.RenderSystem interface by queueing this call so that it
// gets executed inside the locked thread.
func (r *RenderSystem) PostMakeWindow(cfg core.WindowConfig, w *glfw.Window) {
	w.MakeContextCurrent()
	glfw.SwapInterval(cfg.Vsync)

	if err := gl.Init(); err != nil {
		panic(err)
	}

	glog.Info("Checking GL Init status")
	glog.Info("OpenGL version: ", gl.GoStr(gl.GetString(gl.VERSION)))
	glog.Info("OpenGL renderer: ", gl.GoStr(gl.GetString(gl.RENDERER)))

	var (
		maxTextureBufferSize      int32
		maxUniformBlockSize       int32
		maxShaderStorageBlockSize int32
		maxVertexUniformVectors   int32
	)
	gl.GetIntegerv(gl.MAX_TEXTURE_BUFFER_SIZE, &maxTextureBufferSize)
	gl.GetIntegerv(gl.MAX_UNIFORM_BLOCK_SIZE, &maxUniformBlockSize)
	gl.GetIntegerv(gl.MAX_SHADER_STORAGE_BLOCK_SIZE, &maxShaderStorageBlockSize)
	gl.GetIntegerv(gl.MAX_VERTEX_UNIFORM_VECTORS, &maxVertexUniformVectors)

	glog.Info("Maximum buffer sizes:")
	glog.Infof("\tTextureBuffer: %d", maxTextureBufferSize)
	glog.Infof("\tUniformBlock: %d", maxUniformBlockSize)
	glog.Infof("\tShaderStorageBlock (bytes): %d", maxShaderStorageBlockSize)
	glog.Infof("\tUniformVectors: %d", maxVertexUniformVectors)

	// generate basic mesh buffers
	sharedBuffers = newBuffers()
	imguiBuffers = newBuffers()

	// current state
	clearState = core.GetResourceManager().State("clear")
	currentState = clearState
	bindState(clearState, true)
}

// PreMakeWindow is called before the actual glfw window creation to allow
// the rendersystem to provide its hints (API, versions, etc)
func (r *RenderSystem) PreMakeWindow() {
	glog.Info("Setting window context tips, Requesting OpenGL4.1 core")

	// create a window
	glfw.WindowHint(glfw.Decorated, glfw.True)
	glfw.WindowHint(glfw.Visible, glfw.True)
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 3)
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)
	glfw.WindowHint(glfw.DepthBits, 32)
	glfw.WindowHint(glfw.Samples, 0)
}
