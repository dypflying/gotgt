package backingstore

import (
	"github.com/dypflying/go-qcow2lib/qcow2"
	log "github.com/sirupsen/logrus"

	"github.com/gostor/gotgt/pkg/api"
	"github.com/gostor/gotgt/pkg/scsi"
)

// This ceph-rbd plugin is only for linux
// path format ceph-rbd:poolname/imagename
const (
	Qcow2BackingStorage = "qcow2"
)

func init() {
	scsi.RegisterBackingStore(Qcow2BackingStorage, newQcow2)
}

type qcow2Store struct {
	scsi.BaseBackingStore
	child *qcow2.BdrvChild
}

func newQcow2() (api.BackingStore, error) {
	return &qcow2Store{
		BaseBackingStore: scsi.BaseBackingStore{
			Name:            Qcow2BackingStorage,
			DataSize:        0,
			OflagsSupported: 0,
		},
	}, nil
}

func (bs *qcow2Store) Open(dev *api.SCSILu, path string) error {
	var err error
	var open_opts = map[string]any{
		qcow2.OPT_FILENAME: path,
		qcow2.OPT_FMT:      "qcow2",
	}
	log.Infof("open qcow2 path = %s", path)
	if bs.child, err = qcow2.Blk_Open(path, open_opts, qcow2.BDRV_O_RDWR); err != nil {
		return err
	}
	if bs.DataSize, err = qcow2.Blk_Getlength(bs.child); err != nil {
		return err
	}

	return nil
}

func (bs *qcow2Store) Close(dev *api.SCSILu) error {
	log.Infof("Close qcow2")
	qcow2.Blk_Close(bs.child)
	return nil
}

func (bs *qcow2Store) Init(dev *api.SCSILu, Opts string) error {
	log.Infof("Init qcow2 opts = %s", Opts)
	return nil
}

func (bs *qcow2Store) Exit(dev *api.SCSILu) error {
	log.Infof("Exit qcow2")
	return nil
}

func (bs *qcow2Store) Size(dev *api.SCSILu) uint64 {
	log.Infof("Size qcow2")
	return bs.DataSize
}

func (bs *qcow2Store) Read(offset, tl int64) ([]byte, error) {
	var err error
	tmpbuf := make([]byte, tl)
	log.Debugf("qcow2 read bytes=%d", tl)
	_, err = qcow2.Blk_Pread(bs.child, uint64(offset), tmpbuf, uint64(tl))
	return tmpbuf, err
}

func (bs *qcow2Store) Write(wbuf []byte, offset int64) error {
	var err error
	bytes := uint64(len(wbuf))
	log.Debugf("qcow2 write bytes=%d", bytes)
	_, err = qcow2.Blk_Pwrite(bs.child, uint64(offset), wbuf, bytes, 0)
	return err
}

func (bs *qcow2Store) DataSync(offset, tl int64) error {
	log.Infof("DataSync qcow2 offset=%d, tl=%d", offset, tl)
	qcow2.Blk_Flush(bs.child)
	return nil
}

func (bs *qcow2Store) DataAdvise(offset, length int64, advise uint32) error {
	log.Infof("DataAdvise qcow2 offset=%d, length=%d, advise=%d", offset, length, advise)
	return nil
}

func (bs *qcow2Store) Unmap(descs []api.UnmapBlockDescriptor) error {
	for i := 0; i < len(descs); i++ {
		log.Infof("Unmap qcow2 offset=%d, tl=%d", descs[i].Offset, descs[i].TL)
		qcow2.Blk_Discard(bs.child, descs[i].Offset, uint64(descs[i].TL))
	}
	return nil
}
