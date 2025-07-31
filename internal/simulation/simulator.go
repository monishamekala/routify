package simulation

import (
    "math/rand"
    "time"
)

// TrafficCondition represents a real-time traffic modifier
type TrafficCondition struct {
    SegmentID string
    Delay     float32 // delay in minutes
    Severity  string  // "low", "medium", "high"
}

// SimulateTraffic returns simulated traffic delays for route segments
func SimulateTraffic(segmentIDs []string) []TrafficCondition {
    rand.Seed(time.Now().UnixNano())
    var results []TrafficCondition

    for _, id := range segmentIDs {
        delay := rand.Float32() * 5 // simulate up to 5 minutes of delay
        severity := "low"
        if delay > 4 {
            severity = "high"
        } else if delay > 2 {
            severity = "medium"
        }

        results = append(results, TrafficCondition{
            SegmentID: id,
            Delay:     delay,
            Severity:  severity,
        })
    }

    return results
}
