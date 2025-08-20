# Simple Factory Pattern

## Amaç
- Object oluşturma sürecini merkezi bir noktaya taşımak.
- Client’ların doğrudan `new` ile concrete instance yaratmasını engellemek.

## Kullanım Alanları
- Birden fazla concrete type'ın tek bir interface üzerinden yaratıldığı durumlar.
- Nesnelerin oluşturulma mantığı karmaşık veya koşullara bağlıysa.
- Factory pattern’in basit versiyonudur (abstract factory’e göre daha az esnektir).