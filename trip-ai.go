package model

// Thời gian di chuyển
type Trip struct {
    MaChuyenXe  string `json:"maChuyenXe"`
    TenChuyenXe string `json:"tenChuyenXe"`
    TuyenDuong  string `json:"tuyenDuong"`
    NgayDi      string `json:"ngayDi"`
    GioDi       string `json:"gioDi"`
    ThoiGianDi  string `json:"thoiGianDi"` // Ví dụ: 5h30m
}

// Tạo thời gian di chuyển mới
func NewTrip(maChuyenXe, tenChuyenXe, tuyenDuong, ngayDi, gioDi, thoiGianDi string) Trip {
    return Trip{
        MaChuyenXe:  maChuyenXe,
        TenChuyenXe: tenChuyenXe,
        TuyenDuong:  tuyenDuong,
        NgayDi:      ngayDi,
        GioDi:       gioDi,
        ThoiGianDi:  thoiGianDi,
    }
}

// Chỉnh sửa thời gian di chuyển
func (t *Trip) Edit(maChuyenXe, tenChuyenXe, tuyenDuong, ngayDi, gioDi, thoiGianDi string) {
    t.MaChuyenXe = maChuyenXe
    t.TenChuyenXe = tenChuyenXe
    t.TuyenDuong = tuyenDuong
    t.NgayDi = ngayDi
    t.GioDi = gioDi
    t.ThoiGianDi = thoiGianDi
}
