APP_NAME := estuge
VERSION := 1.1.1
BINARY := $(APP_NAME)
MAINTAINER := Rick Collette <megalith@root.sh>
DESCRIPTION := Lightweight markdown to HTML static website generator
DEB_PACKAGE := $(APP_NAME)_$(VERSION)_amd64.deb
RPM_PACKAGE := $(APP_NAME)-$(VERSION)-1.x86_64.rpm

build: structure
	go build -o $(BINARY) main.go

structure:
	mkdir -p usr/local/bin
	mkdir -p usr/share/man/man1

install: build
	install -d /usr/local/bin
	install -m 755 $(BINARY) /usr/local/bin/$(BINARY)
	install -d /usr/share/man/man1
	install -m 644 $(APP_NAME).1 /usr/share/man/man1/$(APP_NAME).1

deb: build
	# Create the directory structure
	mkdir -p debian/DEBIAN
	mkdir -p debian/usr/local/bin
	mkdir -p debian/usr/share/man/man1
	mkdir -p debian/usr/share/doc/$(APP_NAME)

	# Create the control file
	echo "Package: $(APP_NAME)" > debian/DEBIAN/control
	echo "Version: $(VERSION)" >> debian/DEBIAN/control
	echo "Section: utils" >> debian/DEBIAN/control
	echo "Priority: optional" >> debian/DEBIAN/control
	echo "Architecture: amd64" >> debian/DEBIAN/control
	echo "Depends: " >> debian/DEBIAN/control
	echo "Maintainer: $(MAINTAINER)" >> debian/DEBIAN/control
	echo "Description: $(DESCRIPTION)" >> debian/DEBIAN/control

	# Copy files
	install -m 755 $(BINARY) debian/usr/local/bin/
	install -m 644 $(APP_NAME).1 debian/usr/share/man/man1/
	install -m 644 LICENSE debian/usr/share/doc/$(APP_NAME)/

	# Build the deb package
	dpkg-deb --build debian
	mv debian.deb $(DEB_PACKAGE)

rpm: build
	mkdir -p rpm_build/usr/local/bin
	mkdir -p rpm_build/usr/share/man/man1
	mkdir -p rpm_build/usr/share/doc/$(APP_NAME)
	install -m 755 $(BINARY) rpm_build/usr/local/bin/
	install -m 644 $(APP_NAME).1 rpm_build/usr/share/man/man1/
	install -m 644 LICENSE rpm_build/usr/share/doc/$(APP_NAME)/
	if [ -f "rpm_build/usr/local/bin/$(BINARY)" ] && [ -f "rpm_build/usr/share/man/man1/$(APP_NAME).1" ] && [ -f "rpm_build/usr/share/doc/$(APP_NAME)/LICENSE" ]; then \
		fpm -s dir -t rpm -n $(APP_NAME) -v $(VERSION) -C rpm_build \
			--maintainer "$(MAINTAINER)" \
			--description "$(DESCRIPTION)" \
			usr/local/bin/$(BINARY) \
			usr/share/man/man1/$(APP_NAME).1 \
			usr/share/doc/$(APP_NAME)/LICENSE; \
	fi
	mv $(APP_NAME)-$(VERSION)-1.x86_64.rpm tmp_$(RPM_PACKAGE)
	mv tmp_$(RPM_PACKAGE) $(RPM_PACKAGE)
	rm -rf rpm_build

all: build deb rpm

clean:
	rm -f $(BINARY)
	rm -rf debian
	rm -f $(DEB_PACKAGE)
	rm -f $(RPM_PACKAGE)
	rm -rf rpm_build
	rm -rf usr
	
