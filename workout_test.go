package main
    
import  (
    "testing"
)

func TestSteps(t *testing.T) {
    // fn := Steps()
    // if val := fn(); val != 0 {
    //     t.Fatalf("TestSteps does not initialize to 0")
    // }
    a := make([]string, 3)
    a[0] = "Test"

    if a[0] != "Test" {
        t.Fatalf("first element not Test")
    }
}
