package iostreams

import (
	"fmt"
	"testing"
)

func TestStopAlternateScreenBuffer(t *testing.T) {
	ios, _, stdout, _ := Test()
	ios.SetAlternateScreenBufferEnabled(true)

	ios.StartAlternateScreenBuffer()
	fmt.Fprint(ios.Out, "test")
	ios.StopAlternateScreenBuffer()

	// Stopping a subsequent time should no-op.
	ios.StopAlternateScreenBuffer()

	const want = "\x1b[?1049htest\x1b[?1049l"
	if got := stdout.String(); got != want {
		t.Errorf("after IOStreams.StopAlternateScreenBuffer() got %q, want %q", got, want)
	}
}
