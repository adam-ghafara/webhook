# Webhook

Ini adalah Webhook yang digunakan pada web GIS Adam.

## Installment

Silahkan buat Cloud Function di GCP.

Setting API dengan setting berikut:

![Gambar 1](./image/Screenshot%202023-11-06%20100948.png)

Isikan Secret dengan kode sesuka hati.

Isikan Token dengan token yang di generate di website [wa.my.id](wa.my.id)

Next.

Pilih Runtime Go 1.21

Copy Codingan yang ada pada function.go kedalam function.go yang disediakan google.

Berikut codingan : [function.go](./gcf/function.go)

Pastikan samakan Entry Point dengan functions.HTTP pada func init().

![Gambar 2](./image/Screenshot%202023-11-06%20101804.png)

Deploy dan lihat hasilnya. Selanjutnya silahkan dicoba di [APIDocs](wa.my.id/apidocs).