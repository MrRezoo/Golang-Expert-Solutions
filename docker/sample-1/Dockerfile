# Dockerfile
# استفاده از یک ایمیج پایه از Go
FROM golang:1.16

# تنظیم دایرکتوری کار
WORKDIR /app

# کپی فایل‌های پروژه به دایرکتوری کار
COPY . .

# کامپایل برنامه
RUN go build -o main .

# تعیین پورت برای اپلیکیشن
EXPOSE 8080

# اجرای برنامه
CMD ["./main"]
