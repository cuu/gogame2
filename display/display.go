package display

import (
  "fmt"
	"os"
	
	"github.com/veandco/go-sdl2/sdl"
	"github.com/cuu/gogame2"
	"github.com/cuu/gogame2/surface"
	
)

var Inited =  false
var window *sdl.Window
var big_surface *sdl.Surface
var renderer *sdl.Renderer
var texture *sdl.Texture
var big_surface_pixels []byte

func AssertInited() {
	if Inited == false {
		panic("run gogame.DisplayInit first")
	}
}

func Init() bool {
	
	sdl.Do(func() {
		
		if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
			panic(err)
		}
	
		Inited = true
	})
	
  return Inited 
}

func Destroy() {
  if window != nil {
    window.Destroy()
  }
}

func GetWindow() *sdl.Window {
    
    return window
}

func SetWindowPos(win*sdl.Window, x,y int) {
    win.SetPosition(int32(x), int32(y))
}

func GetWindowPos(win*sdl.Window) (int,int) {
  x,y := win.GetPosition()
        
  return int(x),int(y)
}

func SetWindowTitle(win*sdl.Window, tit string) {
    win.SetTitle(tit)
}

func SetWindowOpacity(win*sdl.Window, op float64) {
    win.SetWindowOpacity(float32(op))
}

func SetWindowBordered(win*sdl.Window, b bool) {
    win.SetBordered(b)
}

func SetX11WindowOnTop() {
    
}

func GetCurrentMode( scr_index int) (mode sdl.DisplayMode, err error) {
  
  return sdl.GetCurrentDisplayMode(scr_index)

}

func SetMode(w,h,flags,depth int32) *sdl.Surface {
	var err error
	AssertInited()
	
	sdl.Do(func() {
		video_centered := os.Getenv("SDL_VIDEO_CENTERED")
    if flags & gogame.FIRSTHIDDEN > 0{
			window, err = sdl.CreateWindow("gogame2", -w, -h,w, h,uint32(gogame.SHOWN |(flags&(^gogame.FIRSTHIDDEN))))
            window.SetGrab(false)
        }else {
            if video_centered == "1" {
                window, err = sdl.CreateWindow("gogame2", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
                    w, h, uint32( gogame.SHOWN | flags))
            }else {
                window, err = sdl.CreateWindow("gogame2", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
                    w, h, uint32( gogame.SHOWN | flags))
            }
        }
		
		if err != nil {
			panic(err)
		}
    
    
    renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
    if err != nil {
        fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
        return
    }
    
    big_surface = surface.Surface(int(w),int(h))
    texture,err =  renderer.CreateTextureFromSurface(big_surface)
    if err != nil {
      panic(err)
    }
    
    big_surface_pixels = big_surface.Pixels()
    
	})

	return big_surface
}

func UpdatePixels() {
  sdl.Do(func() {
    texture.Update(nil,big_surface_pixels,int(big_surface.Pitch))
  })
}

func Flip() {
	sdl.Do(func() {
    //texture.Update(nil,big_surface_pixels,int(big_surface.Pitch))
    renderer.Clear()
    renderer.Copy(texture, nil,nil)
    renderer.Present()
	})
}
		


