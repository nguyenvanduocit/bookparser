# Bookparser

một vài script hỗ trợ việc parser sách scan thành ebook một cách dễ dàng hơn.

## Chuẩn bị

Scan sách, cắt hết header, footer, chỉ để lại nội dung. Đánh tên của sách theo đúng thứ tự trang. và lưu vào thư mục images.

Tạo project trên google cloud, enable vision api, tạo service account và tạo key rồi lưu lại với tên: service-account.js

## Parse

Chạy command parse.

Command này sẽ lần lượt gọi Vision API để lấy nội dung của các file ảnh rồi lưu lại vào thư mục text.

Vì mình chỉ muốn làm nhanh chóng, nên hardcode trong file hơi nhiều, mục đích là để cố gắng format lại nội dung text, ví dụ như các chỗ xuống dòng.

Tương lai có thể viết thêm command là fmt để thực hiện việc format này một cách tốt hơn

## Merge

Chạy command merge.

Mục đích là để merge mọi file lại với nhau.

Giờ bạn phải tự vào để thêm định dạng cho text.
