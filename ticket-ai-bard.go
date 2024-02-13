package main

import (
    "fmt"
    "strings"

    "github.com/example/go-ticket-manager/model"
    "github.com/example/go-ticket-manager/service"
)

func main() {
    // Khởi tạo các dịch vụ
    ticketService := service.NewTicketService()
    tripService := service.NewTripService()

    // Vòng lặp chương trình
    for {
        // Hiển thị menu
        fmt.Println("----------------------------------")
        fmt.Println("Chương trình quản lý vé xe")
        fmt.Println("----------------------------------")
        fmt.Println("1. Quản lý vé xe")
        fmt.Println("2. Quản lý thời gian di chuyển")
        fmt.Println("0. Thoát chương trình")
        fmt.Println("----------------------------------")

        // Lựa chọn chức năng
        var choice string
        fmt.Scanln(&choice)

        switch strings.TrimSpace(choice) {
        case "1":
            // Quản lý vé xe
            fmt.Println("Chức năng quản lý vé xe")
            fmt.Println("--------------------------")
            fmt.Println("1. Thêm mới vé xe")
            fmt.Println("2. Sửa thông tin vé xe")
            fmt.Println("3. Xóa vé xe")
            fmt.Println("4. Tìm kiếm vé xe")
            fmt.Println("0. Quay lại menu chính")
            fmt.Println("--------------------------")

            var subChoice string
            fmt.Scanln(&subChoice)

            switch strings.TrimSpace(subChoice) {
            case "1":
                // Thêm mới vé xe
                ticketService.AddTicket()
            case "2":
                // Sửa thông tin vé xe
                ticketService.EditTicket()
            case "3":
                // Xóa vé xe
                ticketService.DeleteTicket()
            case "4":
                // Tìm kiếm vé xe
                ticketService.SearchTicket()
            case "0":
                // Quay lại menu chính
                continue
            default:
                fmt.Println("Lựa chọn không hợp lệ!")
            }
        case "2":
            // Quản lý thời gian di chuyển
            fmt.Println("Chức năng quản lý thời gian di chuyển")
            fmt.Println("-------------------------------------")
            fmt.Println("1. Thêm mới thời gian di chuyển cho tuyến xe")
            fmt.Println("2. Sửa thời gian di chuyển cho tuyến xe")
            fmt.Println("3. Xóa thời gian di chuyển cho tuyến xe")
            fmt.Println("4. Tìm kiếm thời gian di chuyển cho tuyến xe")
            fmt.Println("0. Quay lại menu chính")
            fmt.Println("-------------------------------------")

            var subChoice string
            fmt.Scanln(&subChoice)

            switch strings.TrimSpace(subChoice) {
            case "1":
                // Thêm mới thời gian di chuyển cho tuyến xe
                tripService.AddTrip()
            case "2":
                // Sửa thời gian di chuyển cho tuyến xe
                tripService.EditTrip()
            case "3":
                // Xóa thời gian di chuyển cho tuyến xe
                tripService.DeleteTrip()
            case "4":
                // Tìm kiếm thời gian di chuyển cho tuyến xe
                tripService.SearchTrip()
            case "0":
                // Quay lại menu chính
                continue
            default:
                fmt.Println("Lựa chọn không hợp lệ!")
            }
        case "0":
            // Thoát chương trình
            fmt.Println("Thoát chương
