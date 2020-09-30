package perfume

import (
	"fmt"
	"testing"
)

func percent(v uint, p float32) float32 {
	return float32(v) * (p / 100)
}
func getLenUint(s string) uint {
	return uint(len(s))
}
func middleRel(width, height uint, localText string) RelLocation {
	return NewRelativeLocation(int(width/2)-len(localText)/2, int(height/2)-1)
}
func TestMain(t *testing.T) {

	dearTxt := "Dear my lover"
	fromTxt := "From your lover"
	content := [4]string{
		"I want you to know that there's no one who can replace you.",
		"Everyday seems like a blessing since I have met you.",
		"I'm so completely in love with you.",
		"So, I just wanted to say I love you.",
	}

	fullSize := NewSize(FullSize, FullSize)

	window, err := NewWindow(fullSize)
	if err != nil {
		fmt.Println(err)
		return
	}

	winSize := window.Size()
	h := NewSize(uint(percent(winSize.Height, 15)), winSize.Width)
	b := NewSize(uint(percent(winSize.Height, 60)), winSize.Width)
	f := NewSize(uint(percent(winSize.Height, 25)), winSize.Width)

	head := NewHead(h, "Head")
	body := NewBody(b, "Body")
	foot := NewFooter(f, "Footer")

	headLayout := NewFreeLayout("HeadLayout", h.Plus(-2))
	bodyLayout := NewStackLayout("BodyLayout", b.Plus(-2), VerticalOrientation, 1)
	footLayout := NewFreeLayout("FootLayout", f.Plus(-2))

	dear := NewText("DearText", dearTxt, NewSize(1, getLenUint(dearTxt)))
	contents := make([]*Text, 0)
	for i, c := range content {
		id := fmt.Sprintf("content%d", i)
		contents = append(contents, NewText(id, c, NewSize(1, getLenUint(c))))
	}
	from := NewText("FromText", fromTxt, NewSize(1, getLenUint(fromTxt)))

	dear.SetLocation(middleRel(head.Size().Width, head.Size().Height, dear.Text()))
	from.SetLocation(middleRel(foot.Size().Width, foot.Size().Height, from.Text()))

	borderOpt := NewOption(BorderOption, "")

	borderOpt.Set("-")
	head.AddOption(borderOpt.Clone())

	borderOpt.Set("=")
	body.AddOption(borderOpt.Clone())

	borderOpt.Set("*")
	foot.AddOption(borderOpt.Clone())

	headLayout.AddChild(dear)
	for _, t := range contents {
		bodyLayout.AddChild(t)
	}
	footLayout.AddChild(from)

	head.AddChild(headLayout)
	body.AddChild(bodyLayout)
	foot.AddChild(footLayout)

	window.Add(body)
	window.Add(head)
	window.Add(foot)

	r := NewRenderer(window)

	for {
		if changed := r.Confirm(); changed {
			r.Clear()
			r.Render()
		}
	}
}
