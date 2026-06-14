// Ensure that during a lease transfer, the current closed timestamp state is packaged or coordinated so the incoming leaseholder is aware of the high-water mark.
func (r *Replica) AdminTransferLease(ctx context.Context, req *roachpb.AdminTransferLeaseRequest) (*roachpb.AdminTransferLeaseResponse, error) {
    // Get the current closed timestamp from the previous leaseholder
    prevClosedTimestamp := r.getClosedTimestamp()

    // Transfer the lease and include the closed timestamp in the transfer command
    resp, err := r.raftGroup.TransferLeader(ctx, req.LeaseHolder.StoreID)
    if err != nil {
        return nil, err
    }

    // Update the new leaseholder with the previous closed timestamp
    r.updateClosedTimestamp(prevClosedTimestamp)

    return resp, nil
}
