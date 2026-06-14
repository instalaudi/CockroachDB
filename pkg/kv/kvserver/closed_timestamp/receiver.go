// Followers must reject reads if they detect a lease epoch change until they receive a valid closed timestamp update from the *new* leaseholder.
type Receiver struct {
    // ... other fields
    currentLeaseEpoch int64
}

func (r *Receiver) OnNewLease(lease roachpb.Lease) {
    if r.currentLeaseEpoch != lease.Epoch {
        r.rejectReadsUntilClosedTimestampUpdate()
    }
}
