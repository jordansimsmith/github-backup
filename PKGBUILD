# Maintainer: Jordan Sim-Smith <jordansimsmith@gmail.com>
pkgname=github-backup-git
pkgver=r11.4e84cc2
pkgrel=1
pkgdesc="Simple tool that maintains backups of your GitHub repositories"
arch=("x86_64")
url="https://github.com/jordansimsmith/github-backup.git"
license=("MIT")
groups=()
depends=()
makedepends=("git" "go")
provides=("${pkgname%-git}")
conflicts=("${pkgname%-git}")
replaces=()
backup=()
options=()
install=
source=("git+$url")
noextract=()
md5sums=("SKIP")

pkgver() {
  cd "$srcdir/${pkgname%-git}"
  printf "r%s.%s" "$(git rev-list --count HEAD)" "$(git rev-parse --short HEAD)"
}

prepare() {
  cd "$srcdir/${pkgname%-git}"
  mkdir -p build/
}

build() {
  cd "$srcdir/${pkgname%-git}"
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw"
  go build -o build ./cmd/...
}

check() {
  cd "$srcdir/${pkgname%-git}"
  go test ./...
}

package() {
  cd "$srcdir/${pkgname%-git}"
  install -Dm755 build/${pkgname%-git} "$pkgdir/usr/bin/${pkgname%-git}"
  install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"
  install -Dm644 README.md "$pkgdir/usr/share/doc/$pkgname/README.md"
}
