// Copyright 2012 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package declarative

import (
	"github.com/StevenZack/livedata"
	"github.com/gofaith/walk"
)

type PushButton struct {
	// Window

	Accessibility      Accessibility
	Background         Brush
	ContextMenuItems   []MenuItem
	DoubleBuffering    bool
	Enabled            Property
	Font               Font
	MaxSize            Size
	MinSize            Size
	Name               string
	OnBoundsChanged    walk.EventHandler
	OnKeyDown          walk.KeyEventHandler
	OnKeyPress         walk.KeyEventHandler
	OnKeyUp            walk.KeyEventHandler
	OnMouseDown        walk.MouseEventHandler
	OnMouseMove        walk.MouseEventHandler
	OnMouseUp          walk.MouseEventHandler
	OnSizeChanged      walk.EventHandler
	Persistent         bool
	RightToLeftReading bool
	ToolTipText        Property
	Visible            Property

	// Widget

	Alignment          Alignment2D
	AlwaysConsumeSpace bool
	Column             int
	ColumnSpan         int
	GraphicsEffects    []walk.WidgetGraphicsEffect
	Row                int
	RowSpan            int
	StretchFactor      int

	// Button

	Image     Property
	OnClicked walk.EventHandler
	Text      Property

	// PushButton

	AssignTo       **walk.PushButton
	ImageAboveText bool

	BindVisible *livedata.Bool
	BindEnable  *livedata.Bool
}

func (pb PushButton) Create(builder *Builder) error {
	w, err := walk.NewPushButton(builder.Parent())
	if err != nil {
		return err
	}

	if pb.AssignTo != nil {
		*pb.AssignTo = w
	}

	if pb.BindVisible != nil {
		pb.BindVisible.ObserveForever(func(b bool) {
			w.SetVisible(b)
		})
	}

	if pb.BindEnable != nil {
		pb.BindEnable.ObserveForever(func(b bool) {
			w.SetEnabled(b)
		})
	}
	return builder.InitWidget(pb, w, func() error {
		if err := w.SetImageAboveText(pb.ImageAboveText); err != nil {
			return err
		}

		if pb.OnClicked != nil {
			w.Clicked().Attach(pb.OnClicked)
		}

		return nil
	})
}
