package model

// Vé xe
type Ticket struct {
    ID           string `json:"id"`
    MaChuyenXe  string `json:"maChuyenXe"`
    TenChuyenXe string `json:"tenChuyenXe"`
    TuyenDuong  string `json:"tuyenDuong"`
    TenKhach    string `json:"tenKhach"`
    SDT         string `json:"sdt"`
    NgayDi      string `json:"ngayDi"`
    GioDi       string `json:"gioDi"`
    GiaVe       float64 `json:"giaVe"`
}

// Tạo vé xe mới
func NewTicket(id, maChuyenXe, tenChuyenXe, tuyenDuong, tenKhach, sdt, ngayDi, gioDi string, giaVe float64) Ticket {
    return Ticket{
        ID:           id,
        MaChuyenXe:  maChuyenXe,
        TenChuyenXe: tenChuyenXe,
        TuyenDuong:  tuyenDuong,
        TenKhach:    tenKhach,
        SDT:         sdt,
        NgayDi:      ngayDi,
        GioDi:       gioDi,
        GiaVe:       giaVe,
    }
}

// Chỉnh sửa thông tin vé xe
func (t *Ticket) Edit(maChuyenXe, tenChuyenXe, tuyenDuong, tenKhach, sdt, ngayDi, gioDi string, giaVe float64) {
    t.MaChuyenXe = maChuyenXe
    t.TenChuyenXe = tenChuyenXe
    t.TuyenDuong = tuyenDuong
    t.TenKhach = tenKhach
    t.SDT = sdt
    t.NgayDi = ngayDi
    t.GioDi = gioDi
    t.GiaVe = giaVe
}
