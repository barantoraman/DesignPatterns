# Abstract Factory Pattern

## Amaç
- Birbiriyle ilişkili veya bağımlı nesne gruplarını tek bir interface üzerinden oluşturmak.
- Client’ın concrete sınıflara doğrudan bağımlılığını azaltmak.
- Bir ürün ailesi içindeki nesnelerin birlikte tutarlılığını sağlamak.

## Kullanım Alanları
- Birden fazla nesne ailesi (product family) var ve client bu nesneleri birlikte kullanıyor.
- Oluşturulan nesnelerin uyumlu olması gerekiyor (örn. MacOS widget set vs Windows widget set).
- Factory Method, tek bir nesne üretirken; Abstract Factory bir grup ilişkili nesneyi aynı anda üretebilir.
- UI toolkit, tema sistemi, cross-platform adapter veya plugin sistemlerinde sık kullanılır.