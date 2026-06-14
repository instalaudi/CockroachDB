// Ensure that when a replica loses its lease, it stops publishing closed timestamps,
// and the new leaseholder initializes its closed timestamp queue starting *at least* at the last closed timestamp of the previous lease.
type Tracker struct {
    // ... other fields
    lastClosedTimestamp hlc.Timestamp
}

func (t *Tracker) OnLeaseTransfer(newLeaseHolder roachpb.StoreID) {
    t.lastClosedTimestamp = t.getClosedTimestamp()
    t.stopPublishingClosedTimestamps()
}

func (t *Tracker) InitializeForNewLeaseholder() {
    t.startPublishingClosedTimestamps(t.lastClosedTimestamp)
}
