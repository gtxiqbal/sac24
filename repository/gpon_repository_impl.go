package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/gtxiqbal/sac24/helper"
	"github.com/gtxiqbal/sac24/model/domain"
	"strings"
)

type GponRepositoryImpl struct {
}

func NewGponRepositoryImpl() GponRepository {
	return &GponRepositoryImpl{}
}

func (repository *GponRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.Gpon {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gpons []domain.Gpon
	for rows.Next() {
		var gponUsername sql.NullString
		var gponPassword sql.NullString
		var gponUpdateAt sql.NullTime
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		gpon := domain.Gpon{}

		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		gpons = append(gpons, gpon)
	}

	return gpons
}

func (repository *GponRepositoryImpl) FindById(ctx context.Context, db *sql.DB, id string) (domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND g.id = $1
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		id,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gponUsername sql.NullString
	var gponPassword sql.NullString
	var gponUpdateAt sql.NullTime
	var nmsUsername sql.NullString
	var nmsPassword sql.NullString
	var nmsUpdateAt sql.NullTime
	var stoUpdateAt sql.NullTime
	var witelUpdateAt sql.NullTime
	var regionalUpdateAt sql.NullTime
	gpon := domain.Gpon{}

	if rows.Next() {
		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time

		return gpon, nil
	}

	return gpon, errors.New("gpon not found by id")
}

func (repository *GponRepositoryImpl) FindByHostname(ctx context.Context, db *sql.DB, hostname string) (domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND g.hostname = $1
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		hostname,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gponUsername sql.NullString
	var gponPassword sql.NullString
	var gponUpdateAt sql.NullTime
	var nmsUsername sql.NullString
	var nmsPassword sql.NullString
	var nmsUpdateAt sql.NullTime
	var stoUpdateAt sql.NullTime
	var witelUpdateAt sql.NullTime
	var regionalUpdateAt sql.NullTime
	gpon := domain.Gpon{}

	if rows.Next() {
		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time

		return gpon, nil
	}

	return gpon, errors.New("gpon not found by hostname")
}

func (repository *GponRepositoryImpl) FindByIpAddress(ctx context.Context, db *sql.DB, ipAddress string) (domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND g.ip_address = $1
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		ipAddress,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gponUsername sql.NullString
	var gponPassword sql.NullString
	var gponUpdateAt sql.NullTime
	var nmsUsername sql.NullString
	var nmsPassword sql.NullString
	var nmsUpdateAt sql.NullTime
	var stoUpdateAt sql.NullTime
	var witelUpdateAt sql.NullTime
	var regionalUpdateAt sql.NullTime
	gpon := domain.Gpon{}

	if rows.Next() {
		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time

		return gpon, nil
	}

	return gpon, errors.New("gpon not found by ip address")
}

func (repository *GponRepositoryImpl) FindByIpAddressIn(ctx context.Context, db *sql.DB, ipAddresses []string) ([]domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND g.ip_address = ANY($1::VARCHAR[])
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		"{"+strings.Join(ipAddresses, ",")+"}",
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gpons []domain.Gpon
	for rows.Next() {
		var gponUsername sql.NullString
		var gponPassword sql.NullString
		var gponUpdateAt sql.NullTime
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		gpon := domain.Gpon{}

		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time

		gpons = append(gpons, gpon)
	}

	if len(gpons) > 0 {
		return gpons, nil
	}
	return gpons, errors.New("gpon not found by ip address")
}

func (repository *GponRepositoryImpl) FindByIpAddressInAndProtocol(ctx context.Context, db *sql.DB, ipAddresses []string, protocol string) ([]domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND g.ip_address = ANY($1::VARCHAR[])
				  AND n.protocol = $2
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		"{"+strings.Join(ipAddresses, ",")+"}", protocol,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gpons []domain.Gpon
	for rows.Next() {
		var gponUsername sql.NullString
		var gponPassword sql.NullString
		var gponUpdateAt sql.NullTime
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		gpon := domain.Gpon{}

		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time

		gpons = append(gpons, gpon)
	}

	if len(gpons) > 0 {
		return gpons, nil
	}
	return gpons, errors.New("gpon not found by ip address and protocol")
}

func (repository *GponRepositoryImpl) FindByNmsId(ctx context.Context, db *sql.DB, nmsId string) ([]domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND n.id = $1
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		nmsId,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gpons []domain.Gpon
	for rows.Next() {
		var gponUsername sql.NullString
		var gponPassword sql.NullString
		var gponUpdateAt sql.NullTime
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		gpon := domain.Gpon{}

		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		gpons = append(gpons, gpon)
	}

	if len(gpons) > 0 {
		return gpons, nil
	}
	return gpons, errors.New("gpon not found by nms id")
}

func (repository *GponRepositoryImpl) FindByNmsIpServer(ctx context.Context, db *sql.DB, nmsIpServer string) ([]domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND n.ip_server = $1
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		nmsIpServer,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gpons []domain.Gpon
	for rows.Next() {
		var gponUsername sql.NullString
		var gponPassword sql.NullString
		var gponUpdateAt sql.NullTime
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		gpon := domain.Gpon{}

		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		gpons = append(gpons, gpon)
	}

	if len(gpons) > 0 {
		return gpons, nil
	}
	return gpons, errors.New("gpon not found by nms ip server")
}

func (repository *GponRepositoryImpl) FindByStoId(ctx context.Context, db *sql.DB, stoId string) ([]domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND s.id = $1
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		stoId,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gpons []domain.Gpon
	for rows.Next() {
		var gponUsername sql.NullString
		var gponPassword sql.NullString
		var gponUpdateAt sql.NullTime
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		gpon := domain.Gpon{}

		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		gpons = append(gpons, gpon)
	}

	if len(gpons) > 0 {
		return gpons, nil
	}
	return gpons, errors.New("gpon not found by sto id")
}

func (repository *GponRepositoryImpl) FindByWitelId(ctx context.Context, db *sql.DB, witelId string) ([]domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND w.id = $1
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		witelId,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gpons []domain.Gpon
	for rows.Next() {
		var gponUsername sql.NullString
		var gponPassword sql.NullString
		var gponUpdateAt sql.NullTime
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		gpon := domain.Gpon{}

		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		gpons = append(gpons, gpon)
	}

	if len(gpons) > 0 {
		return gpons, nil
	}
	return gpons, errors.New("gpon not found by witel id")
}

func (repository *GponRepositoryImpl) FindByRegionalId(ctx context.Context, db *sql.DB, regionalId string) ([]domain.Gpon, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT g.id, g.hostname, g.ip_address, g.vlan_inet, g.vlan_voice, g.default_user_nms, g.username, g."password", g.create_at, g.update_at, 
    				n.id, n.nama, n.vendor, n.ip_server, n.port_tl1, n.protocol, n.username, n."password", n.create_at, n.update_at,
       				s.id, s.alias, s.create_at, s.nama, s.update_at,
       				w.id, w.alias, w.create_at, w.nama, w.update_at,
    				r.id, r.create_at, r.nama, r.update_at 
				FROM gpon g, nms n, sto s, witel w, regional r 
				WHERE g.nms_id = n.id
				  AND g.sto_id = s.id
				  AND s.witel_id = w.id
				  AND w.regional_id = r.id 
				  AND r.id = $1
				ORDER BY r.id, w.alias, s.alias, g.hostname `,
		regionalId,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var gpons []domain.Gpon
	for rows.Next() {
		var gponUsername sql.NullString
		var gponPassword sql.NullString
		var gponUpdateAt sql.NullTime
		var nmsUsername sql.NullString
		var nmsPassword sql.NullString
		var nmsUpdateAt sql.NullTime
		var stoUpdateAt sql.NullTime
		var witelUpdateAt sql.NullTime
		var regionalUpdateAt sql.NullTime
		gpon := domain.Gpon{}

		err = rows.Scan(
			&gpon.Id,
			&gpon.Hostname,
			&gpon.IpAddress,
			&gpon.VlanInet,
			&gpon.VlanVoice,
			&gpon.DefaultUserNms,
			&gponUsername,
			&gponPassword,
			&gpon.CreateAt,
			&gponUpdateAt,
			&gpon.Nms.Id,
			&gpon.Nms.Nama,
			&gpon.Nms.Vendor,
			&gpon.Nms.IpServer,
			&gpon.Nms.PortTl1,
			&gpon.Nms.Protocol,
			&nmsUsername,
			&nmsPassword,
			&gpon.Nms.CreateAt,
			&nmsUpdateAt,
			&gpon.Sto.Id,
			&gpon.Sto.Alias,
			&gpon.Sto.CreateAt,
			&gpon.Sto.Nama,
			&stoUpdateAt,
			&gpon.Sto.Witel.Id,
			&gpon.Sto.Witel.Alias,
			&gpon.Sto.Witel.CreateAt,
			&gpon.Sto.Witel.Nama,
			&witelUpdateAt,
			&gpon.Sto.Witel.Regional.Id,
			&gpon.Sto.Witel.Regional.CreateAt,
			&gpon.Sto.Witel.Regional.Nama,
			&regionalUpdateAt,
		)
		helper.PanicIfError(err)

		gpon.Username = gponUsername.String
		gpon.Password = gponPassword.String
		gpon.UpdateAt = gponUpdateAt.Time
		gpon.Nms.Username = nmsUsername.String
		gpon.Nms.Password = nmsPassword.String
		gpon.Nms.UpdateAt = nmsUpdateAt.Time
		gpon.Sto.UpdateAt = stoUpdateAt.Time
		gpon.Sto.Witel.UpdateAt = witelUpdateAt.Time
		gpon.Sto.Witel.Regional.UpdateAt = regionalUpdateAt.Time
		gpons = append(gpons, gpon)
	}

	if len(gpons) > 0 {
		return gpons, nil
	}
	return gpons, errors.New("gpon not found by regional id")
}

func (repository *GponRepositoryImpl) FindVendorByIpAddress(ctx context.Context, db *sql.DB, ipAddress string) (string, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT DISTINCT n.vendor 
			   FROM gpon g, nms n 
			   WHERE g.nms_id = n.id 
			     AND g.ip_address = $1`,
		ipAddress,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var vendor string
	if rows.Next() {
		err = rows.Scan(&vendor)
		helper.PanicIfError(err)

		return vendor, nil
	}
	return "", errors.New("vendor not found by ip address")
}

func (repository *GponRepositoryImpl) FindVendorByIpAddressIn(ctx context.Context, db *sql.DB, ipAddresses []string) ([]string, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT DISTINCT n.vendor 
			   FROM gpon g, nms n 
			   WHERE g.nms_id = n.id 
			     AND g.ip_address = ANY($1::VARCHAR[])`,
		"{"+strings.Join(ipAddresses, ",")+"}")
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var vendors []string
	for rows.Next() {
		var vendor string
		err = rows.Scan(&vendor)
		helper.PanicIfError(err)

		vendors = append(vendors, vendor)
	}
	if len(vendors) > 0 {
		return vendors, nil
	}
	return vendors, errors.New("vendor not found by list ip address")
}

func (repository *GponRepositoryImpl) FindProtocolByIpAddress(ctx context.Context, db *sql.DB, ipAddress string) (string, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT DISTINCT n.protocol 
			   FROM gpon g, nms n 
			   WHERE g.nms_id = n.id 
			     AND g.ip_address = $1`,
		ipAddress,
	)
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var vendor string
	if rows.Next() {
		err = rows.Scan(&vendor)
		helper.PanicIfError(err)

		return vendor, nil
	}
	return "", errors.New("protocol not found by ip address")
}

func (repository *GponRepositoryImpl) FindProtocolByIpAddressIn(ctx context.Context, db *sql.DB, ipAddresses []string) ([]string, error) {
	rows, err := db.QueryContext(ctx,
		`SELECT DISTINCT n.protocol 
			   FROM gpon g, nms n 
			   WHERE g.nms_id = n.id 
			     AND g.ip_address = ANY($1::VARCHAR[])`,
		"{"+strings.Join(ipAddresses, ",")+"}")
	helper.PanicIfError(err)
	defer helper.RowsClose(rows)

	var protocols []string
	for rows.Next() {
		var protocol string
		err = rows.Scan(&protocol)
		helper.PanicIfError(err)

		protocols = append(protocols, protocol)
	}
	if len(protocols) > 0 {
		return protocols, nil
	}
	return protocols, errors.New("protocol not found by list ip address")
}
