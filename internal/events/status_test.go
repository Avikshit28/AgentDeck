package events
import "testing"
func TestStatusString (t *testing.T){
	tests := []struct {
		status Status
		want string
	}{
		{StatusPending, "pending"},
		{StatusRunning, "running"},
		{StatusWaiting, "waiting"},
		{StatusDone, "done"},
		{StatusFailed, "failed"},
		{Status(99), "unknown"},
	}
	for _, tt := range tests {
		got := tt.status.String()
		if got != tt.want{
			t.Errorf("Status(%d).String() = %q, want %q", int(tt.status), got, tt.want)
		}
	}
}
func TestStatusTerminal (t *testing.T){
	if StatusRunning.Terminal(){
		t.Error("running should not be in terminal")
	}
	if !StatusDone.Terminal(){
		t.Error("done should be terminal")
	}
	if !StatusFailed.Terminal(){
		t.Error("failed should be terminal")
	}
}