package managers

import (
	"github.com/hajimehoshi/ebiten/v2"
)

var KeypressManager *Keypress

type keyBindType int

const (
	keyBindPress        keyBindType = 0
	keyBindPressAndHold keyBindType = 1
)

type boundKey struct {
	bindType keyBindType
	listener func()
}

type Keypress struct {
	listeners map[string]map[string]boundKey
	pressed   map[string]bool
}

func (kp *Keypress) OnUpdate(keys []ebiten.Key) {
	newPressed := map[string]bool{}
	for _, k := range keys {
		if _, ok := kp.listeners[ebiten.KeyName(k)]; ok {
			newPressed[ebiten.KeyName(k)] = true
			for _, v2 := range kp.listeners[ebiten.KeyName(k)] {
				switch v2.bindType {
				case keyBindPress:
					if _, ok := kp.pressed[ebiten.KeyName(k)]; !ok {
						v2.listener()
					}
				default:
					v2.listener()
				}
			}
		}
	}

	kp.pressed = newPressed
}

func (kp *Keypress) BindPress(id string, key ebiten.Key, listener func()) {
	if _, ok := kp.listeners[ebiten.KeyName(key)]; ok {
		kp.listeners[ebiten.KeyName(key)][id] = boundKey{listener: listener, bindType: keyBindPress}
	} else {
		kp.listeners[ebiten.KeyName(key)] = map[string]boundKey{}
		kp.listeners[ebiten.KeyName(key)][id] = boundKey{listener: listener, bindType: keyBindPress}
	}
}

func (kp *Keypress) Bind(id string, key ebiten.Key, listener func()) {
	if _, ok := kp.listeners[ebiten.KeyName(key)]; ok {
		kp.listeners[ebiten.KeyName(key)][id] = boundKey{listener: listener, bindType: keyBindPressAndHold}
	} else {
		kp.listeners[ebiten.KeyName(key)] = map[string]boundKey{}
		kp.listeners[ebiten.KeyName(key)][id] = boundKey{listener: listener, bindType: keyBindPressAndHold}
	}
}

func (kp *Keypress) UnBind(id string) {
	delete(kp.listeners, id)
}

func NewKeyPressManager() {
	KeypressManager = &Keypress{listeners: map[string]map[string]boundKey{}}
}
