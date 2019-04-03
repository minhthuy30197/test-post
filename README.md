
# Test hiển thị các loại bài viết
--- 

##### Chạy ví dụ
- Blog, tải tài liệu, đồ án,... thì sẽ được lưu chung dưới dạng "post" có type để phân biệt. Các loại này có một số trường chung (ví dụ title, tác giả, mô tả,...) còn các trường khác nhau (ví dụ tải tài liệu có link file tài liệu,...) sẽ được định nghĩa dưới dạng json. Khi lấy dữ liệu từ database, tuỳ vào type sẽ parse json sang định dạng struct thích hợp để trả về.
- Truy cập blog, tài liệu, tài liệ,... thông qua url có dạng **/bai-viet?type=...** (ví dụ: /bai-viet?type=du-an). Dựa vào type trên query, render view template và dữ liệu tương ứng.

##### Chạy ví dụ
Truy cập vào thư mục test-post, chạy lệnh sau đó truy cập [localhost:8080](http://localhost:8080) xem kết quả
```
go run main.go
```