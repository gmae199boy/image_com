# 이미지 조합 프로그램

## 결과
 
 <img width="350" alt="스크린샷 2022-10-31 오전 9 51 37" src="https://user-images.githubusercontent.com/34611716/198913203-4dd87616-5593-49aa-bf1b-16c0aa157552.png">
 
## 문제점

<img width="1196" alt="스크린샷 2022-10-31 오전 9 55 07" src="https://user-images.githubusercontent.com/34611716/198913227-2ded63af-c65c-4b3d-8c56-f28f46caaacc.png">

<img width="1194" alt="스크린샷 2022-10-31 오전 9 55 59" src="https://user-images.githubusercontent.com/34611716/198913238-0e8447d7-1ebd-45cd-af97-c1cbd33baab2.png">

<img width="1191" alt="스크린샷 2022-10-31 오전 9 51 17" src="https://user-images.githubusercontent.com/34611716/198913258-cabc1da1-3d25-4630-9d98-9608986f3fe1.png">

- 메모리 점유율이 점점 증가함

## 해결될거 같은 원인

<img width="476" alt="스크린샷 2022-10-31 오전 10 24 10" src="https://user-images.githubusercontent.com/34611716/198913307-46d6c333-0d77-43e5-9338-b477a4c1ffcd.png">

- []*image.Image{...} excapes to heap 이 원인으로 보임
