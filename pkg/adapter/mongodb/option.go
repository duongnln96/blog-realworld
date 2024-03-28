package mongodb

import (
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type ReadConcern string

const (
	LocalReadConcern        ReadConcern = "local"
	MajorityReadConcern     ReadConcern = "majority"
	LinearizableReadConcern ReadConcern = "linearizable"
	AvailableReadConcern    ReadConcern = "available"
	SnapshotReadConcern     ReadConcern = "snapshot"
)

type WriteConcern string

const (
	JournaledWriteConcern      WriteConcern = "journaled"
	MajorityWriteConcern       WriteConcern = "majority"
	UnacknowledgedWriteConcern WriteConcern = "unacknowledged"
	W1WriteConcern             WriteConcern = "w1"
)

type ReadPref string

const (
	NearestReadPref            ReadPref = "nearest"
	PrimaryReadPref            ReadPref = "primary"
	PrimaryPreferredReadPref   ReadPref = "primary_preferred"
	SecondaryReadPref          ReadPref = "secondary_read"
	SecondaryPreferredReadPref ReadPref = "secondary_preferred"
)

type DatabaseOptions struct {
	options *options.DatabaseOptions
}

func NewDatabaseOptions() *DatabaseOptions {
	dbOpts := options.Database()
	return &DatabaseOptions{
		options: dbOpts,
	}
}

func (m *DatabaseOptions) rawMongoOptions() *options.DatabaseOptions {
	return m.options
}

// SetReadConcern sets the value for the ReadConcern field.
func (d *DatabaseOptions) SetReadConcern(rcStr ReadConcern) {
	var rc *readconcern.ReadConcern

	switch rcStr {
	case LocalReadConcern:
		rc = readconcern.Local()

	case MajorityReadConcern:
		rc = readconcern.Majority()

	case LinearizableReadConcern:
		rc = readconcern.Linearizable()

	case AvailableReadConcern:
		rc = readconcern.Available()

	case SnapshotReadConcern:
		rc = readconcern.Snapshot()
	}

	d.options.SetReadConcern(rc)
}

// SetWriteConcern sets the value for the WriteConcern field.
func (d *DatabaseOptions) SetWriteConcern(wcStr WriteConcern) {

	var wc *writeconcern.WriteConcern

	switch wcStr {
	case JournaledWriteConcern:
		wc = writeconcern.Journaled()

	case MajorityWriteConcern:
		wc = writeconcern.Majority()

	case UnacknowledgedWriteConcern:
		wc = writeconcern.Unacknowledged()

	case W1WriteConcern:
		wc = writeconcern.W1()
	}

	d.options.SetWriteConcern(wc)
}

// SetReadPreference sets the value for the ReadPreference field.
func (d *DatabaseOptions) SetReadPreference(rpStr ReadPref) {

	var rp *readpref.ReadPref

	switch rpStr {
	case NearestReadPref:
		rp = readpref.Nearest()
	case PrimaryReadPref:
		rp = readpref.Primary()

	case PrimaryPreferredReadPref:
		rp = readpref.PrimaryPreferred()

	case SecondaryReadPref:
		rp = readpref.Secondary()

	case SecondaryPreferredReadPref:
		rp = readpref.SecondaryPreferred()

	}

	d.options.SetReadPreference(rp)
}

type CollectionOptions struct {
	options *options.CollectionOptions
}

func NewCollectionOptions() *CollectionOptions {
	collOpts := options.Collection()
	return &CollectionOptions{
		options: collOpts,
	}
}

func (m *CollectionOptions) rawCollectionOptions() *options.CollectionOptions {
	return m.options
}

// SetReadConcern sets the value for the ReadConcern field.
func (d *CollectionOptions) SetReadConcern(rcStr ReadConcern) {
	var rc *readconcern.ReadConcern

	switch rcStr {
	case LocalReadConcern:
		rc = readconcern.Local()

	case MajorityReadConcern:
		rc = readconcern.Majority()

	case LinearizableReadConcern:
		rc = readconcern.Linearizable()

	case AvailableReadConcern:
		rc = readconcern.Available()

	case SnapshotReadConcern:
		rc = readconcern.Snapshot()
	}

	d.options.SetReadConcern(rc)
}

// SetWriteConcern sets the value for the WriteConcern field.
func (d *CollectionOptions) SetWriteConcern(wcStr WriteConcern) {

	var wc *writeconcern.WriteConcern

	switch wcStr {
	case JournaledWriteConcern:
		wc = writeconcern.Journaled()

	case MajorityWriteConcern:
		wc = writeconcern.Majority()

	case UnacknowledgedWriteConcern:
		wc = writeconcern.Unacknowledged()

	case W1WriteConcern:
		wc = writeconcern.W1()
	}

	d.options.SetWriteConcern(wc)
}

// SetReadPreference sets the value for the ReadPreference field.
func (d *CollectionOptions) SetReadPreference(rpStr ReadPref) {

	var rp *readpref.ReadPref

	switch rpStr {
	case NearestReadPref:
		rp = readpref.Nearest()

	case PrimaryReadPref:
		rp = readpref.Primary()

	case PrimaryPreferredReadPref:
		rp = readpref.PrimaryPreferred()

	case SecondaryReadPref:
		rp = readpref.Secondary()

	case SecondaryPreferredReadPref:
		rp = readpref.SecondaryPreferred()

	}

	d.options.SetReadPreference(rp)
}
