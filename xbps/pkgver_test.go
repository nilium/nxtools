package xbps

import "testing"

type testCase interface {
	name() string
	run(*testing.T)
}

// pkgVerSuccCase is a test case for successful parsing of an input string to a desired value.
type pkgVerSuccCase struct {
	In   string
	Want PkgVer
}

func (c pkgVerSuccCase) name() string {
	return c.In
}

func (c pkgVerSuccCase) run(t *testing.T) {
	pkgver, err := ParsePkgVer(c.In)
	if err != nil {
		t.Fatalf("ParsePkgVer(%q) error = %v", c.In, err)
	}

	if pkgver != c.Want {
		t.Fatalf("ParsePkgVer(%q): got %#+v; want %#+v", c.In, pkgver, c.Want)
	}

	if got := pkgver.String(); got != c.In {
		t.Fatalf("%#v.String() = %q; want %q", pkgver, got, c.In)
	}
}

// pkgVerFailCase is a test case for returning an error for a specific input.
type pkgVerFailCase struct {
	In  string
	Err error
}

func (c pkgVerFailCase) name() string {
	return c.In
}

func (c pkgVerFailCase) run(t *testing.T) {
	_, err := ParsePkgVer(c.In)
	if err == nil {
		t.Fatalf("ParsePkgVer(%q) error = nil; want error", c.In)
	}

	pe, ok := err.(*PkgVerError)
	if !ok || pe == nil {
		t.Fatalf("ParsePkgVer(%q) error is of type %T; want %T", c.In, err, &PkgVerError{})
	}

	if c.Err != nil && c.Err != pe.Err {
		t.Fatalf("ParsePkgVer(%q) error = %v; want %v", c.In, pe.Err, c.Err)
	}

	same := &PkgVerError{PkgVer: c.In, Err: c.Err}
	if got, want := err.Error(), same.Error(); got != want {
		t.Fatalf("ParsePkgVer(%q) error = %q; want %q", c.In, got, want)
	}
}

func TestParsePkgVer(t *testing.T) {
	cases := []testCase{
		// Successful cases
		pkgVerSuccCase{"Bear-2.3.13_1", PkgVer{"Bear", "2.3.13", 1}},
		pkgVerSuccCase{"Haru-2.3.0_1", PkgVer{"Haru", "2.3.0", 1}},
		pkgVerSuccCase{"QMPlay2-devel-18.12.26_1", PkgVer{"QMPlay2-devel", "18.12.26", 1}},
		pkgVerSuccCase{"SoapyHackRF-32bit-0.3.3_2", PkgVer{"SoapyHackRF-32bit", "0.3.3", 2}},
		pkgVerSuccCase{"apr-util-ldap-1.6.1_6", PkgVer{"apr-util-ldap", "1.6.1", 6}},
		pkgVerSuccCase{"baobab-3.30.0_1", PkgVer{"baobab", "3.30.0", 1}},
		pkgVerSuccCase{"bat-0.9.0_1", PkgVer{"bat", "0.9.0", 1}},
		pkgVerSuccCase{"bglibs-devel-2.04_1", PkgVer{"bglibs-devel", "2.04", 1}},
		pkgVerSuccCase{"bitlbee-facebook-32bit-1.2.0_1", PkgVer{"bitlbee-facebook-32bit", "1.2.0", 1}},
		pkgVerSuccCase{"boinc-devel-7.14.2_2", PkgVer{"boinc-devel", "7.14.2", 2}},
		pkgVerSuccCase{"bsatool-0.44.0_4", PkgVer{"bsatool", "0.44.0", 4}},
		pkgVerSuccCase{"bum-0.1.3_3", PkgVer{"bum", "0.1.3", 3}},
		pkgVerSuccCase{"cbatticon-gtk3-1.6.8_1", PkgVer{"cbatticon-gtk3", "1.6.8", 1}},
		pkgVerSuccCase{"cgrep-6.6.27_1", PkgVer{"cgrep", "6.6.27", 1}},
		pkgVerSuccCase{"clens-0.7.0_3", PkgVer{"clens", "0.7.0", 3}},
		pkgVerSuccCase{"clipper-devel-32bit-6.4.2_2", PkgVer{"clipper-devel-32bit", "6.4.2", 2}},
		pkgVerSuccCase{"compiz-plugins-main-devel-0.8.16_2", PkgVer{"compiz-plugins-main-devel", "0.8.16", 2}},
		pkgVerSuccCase{"cross-mips-linux-musl-0.29_3", PkgVer{"cross-mips-linux-musl", "0.29", 3}},
		pkgVerSuccCase{"curlpp-32bit-0.8.1_1", PkgVer{"curlpp-32bit", "0.8.1", 1}},
		pkgVerSuccCase{"czmq-devel-4.1.1_1", PkgVer{"czmq-devel", "4.1.1", 1}},
		pkgVerSuccCase{"dmd-bootstrap-32bit-2.069.20180305_2", PkgVer{"dmd-bootstrap-32bit", "2.069.20180305", 2}},
		pkgVerSuccCase{"fcgiwrap-1.1.0_4", PkgVer{"fcgiwrap", "1.1.0", 4}},
		pkgVerSuccCase{"fdm-1.9_8", PkgVer{"fdm", "1.9", 8}},
		pkgVerSuccCase{"feh-3.1.1_1", PkgVer{"feh", "3.1.1", 1}},
		pkgVerSuccCase{"flatbuffers-1.10.0_1", PkgVer{"flatbuffers", "1.10.0", 1}},
		pkgVerSuccCase{"font-hanazono-20170904_1", PkgVer{"font-hanazono", "20170904", 1}},
		pkgVerSuccCase{"freefall-4.20_1", PkgVer{"freefall", "4.20", 1}},
		pkgVerSuccCase{"fwupd-devel-1.2.4_1", PkgVer{"fwupd-devel", "1.2.4", 1}},
		pkgVerSuccCase{"gcr-devel-3.28.1_1", PkgVer{"gcr-devel", "3.28.1", 1}},
		pkgVerSuccCase{"gpsd-32bit-3.18.1_1", PkgVer{"gpsd-32bit", "3.18.1", 1}},
		pkgVerSuccCase{"gtk-vnc-devel-0.9.0_1", PkgVer{"gtk-vnc-devel", "0.9.0", 1}},
		pkgVerSuccCase{"gtkimageview-1.6.4_2", PkgVer{"gtkimageview", "1.6.4", 2}},
		pkgVerSuccCase{"gtkmm-32bit-3.24.0_1", PkgVer{"gtkmm-32bit", "3.24.0", 1}},
		pkgVerSuccCase{"harfbuzz-devel-32bit-2.3.1_1", PkgVer{"harfbuzz-devel-32bit", "2.3.1", 1}},
		pkgVerSuccCase{"helm-0.9.0_2", PkgVer{"helm", "0.9.0", 2}},
		pkgVerSuccCase{"hunspell-it_IT-4.2_1", PkgVer{"hunspell-it_IT", "4.2", 1}},
		pkgVerSuccCase{"icecat-i18n-fr-60.3.0_1", PkgVer{"icecat-i18n-fr", "60.3.0", 1}},
		pkgVerSuccCase{"ifupdown-0.8.35_1", PkgVer{"ifupdown", "0.8.35", 1}},
		pkgVerSuccCase{"iperf-2.0.13_1", PkgVer{"iperf", "2.0.13", 1}},
		pkgVerSuccCase{"jasper-devel-32bit-2.0.14_2", PkgVer{"jasper-devel-32bit", "2.0.14", 2}},
		pkgVerSuccCase{"jbigkit-devel-32bit-2.1_2", PkgVer{"jbigkit-devel-32bit", "2.1", 2}},
		pkgVerSuccCase{"kdevelop-32bit-5.3.1_2", PkgVer{"kdevelop-32bit", "5.3.1", 2}},
		pkgVerSuccCase{"knotifyconfig-devel-32bit-5.54.0_1", PkgVer{"knotifyconfig-devel-32bit", "5.54.0", 1}},
		pkgVerSuccCase{"kreport-3.1.0_1", PkgVer{"kreport", "3.1.0", 1}},
		pkgVerSuccCase{"ktoblzcheck-32bit-1.49_1", PkgVer{"ktoblzcheck-32bit", "1.49", 1}},
		pkgVerSuccCase{"kyua-0.13_2", PkgVer{"kyua", "0.13", 2}},
		pkgVerSuccCase{"lft-3.80_1", PkgVer{"lft", "3.80", 1}},
		pkgVerSuccCase{"libArcus-32bit-3.6.0_1", PkgVer{"libArcus-32bit", "3.6.0", 1}},
		pkgVerSuccCase{"libHX-32bit-3.24_1", PkgVer{"libHX-32bit", "3.24", 1}},
		pkgVerSuccCase{"libXdmcp-32bit-1.1.2_3", PkgVer{"libXdmcp-32bit", "1.1.2", 3}},
		pkgVerSuccCase{"libass-0.14.0_1", PkgVer{"libass", "0.14.0", 1}},
		pkgVerSuccCase{"libcddb-devel-32bit-1.3.2_9", PkgVer{"libcddb-devel-32bit", "1.3.2", 9}},
		pkgVerSuccCase{"libcmuclmtk-0.7_3", PkgVer{"libcmuclmtk", "0.7", 3}},
		pkgVerSuccCase{"libevhtp-32bit-1.2.17_1", PkgVer{"libevhtp-32bit", "1.2.17", 1}},
		pkgVerSuccCase{"libfakekey-devel-32bit-0.3_1", PkgVer{"libfakekey-devel-32bit", "0.3", 1}},
		pkgVerSuccCase{"libffi-3.2.1_5", PkgVer{"libffi", "3.2.1", 5}},
		pkgVerSuccCase{"libfm-gtk-devel-1.3.1_1", PkgVer{"libfm-gtk-devel", "1.3.1", 1}},
		pkgVerSuccCase{"libgsf-devel-32bit-1.14.45_1", PkgVer{"libgsf-devel-32bit", "1.14.45", 1}},
		pkgVerSuccCase{"libiec61883-devel-32bit-1.2.0_2", PkgVer{"libiec61883-devel-32bit", "1.2.0", 2}},
		pkgVerSuccCase{"libiscsi-1.18.0_1", PkgVer{"libiscsi", "1.18.0", 1}},
		pkgVerSuccCase{"libmad-devel-0.15.1b_9", PkgVer{"libmad-devel", "0.15.1b", 9}},
		pkgVerSuccCase{"libmate-control-center-devel-32bit-1.20.4_1", PkgVer{"libmate-control-center-devel-32bit", "1.20.4", 1}},
		pkgVerSuccCase{"libmega-32bit-3.2.7_1", PkgVer{"libmega-32bit", "3.2.7", 1}},
		pkgVerSuccCase{"libmusicbrainz5-32bit-5.1.0_4", PkgVer{"libmusicbrainz5-32bit", "5.1.0", 4}},
		pkgVerSuccCase{"libnl-1.1.4_4", PkgVer{"libnl", "1.1.4", 4}},
		pkgVerSuccCase{"liboil-devel-32bit-0.3.17_6", PkgVer{"liboil-devel-32bit", "0.3.17", 6}},
		pkgVerSuccCase{"libosmgpsmap-32bit-1.1.0_3", PkgVer{"libosmgpsmap-32bit", "1.1.0", 3}},
		pkgVerSuccCase{"libpurple-32bit-2.13.0_3", PkgVer{"libpurple-32bit", "2.13.0", 3}},
		pkgVerSuccCase{"libqhull-devel-2015.2_1", PkgVer{"libqhull-devel", "2015.2", 1}},
		pkgVerSuccCase{"libraptor-32bit-2.0.15_2", PkgVer{"libraptor-32bit", "2.0.15", 2}},
		pkgVerSuccCase{"libreoffice-postgresql-32bit-6.1.4.2_2", PkgVer{"libreoffice-postgresql-32bit", "6.1.4.2", 2}},
		pkgVerSuccCase{"libressl-devel-2.8.3_1", PkgVer{"libressl-devel", "2.8.3", 1}},
		pkgVerSuccCase{"librsvg-utils-2.44.12_1", PkgVer{"librsvg-utils", "2.44.12", 1}},
		pkgVerSuccCase{"libsass-3.5.5_1", PkgVer{"libsass", "3.5.5", 1}},
		pkgVerSuccCase{"libshiboken-python3-32bit-1.2.2_4", PkgVer{"libshiboken-python3-32bit", "1.2.2", 4}},
		pkgVerSuccCase{"libspectrum-32bit-1.4.4_1", PkgVer{"libspectrum-32bit", "1.4.4", 1}},
		pkgVerSuccCase{"libtotem-3.30.0_1", PkgVer{"libtotem", "3.30.0", 1}},
		pkgVerSuccCase{"libxmp-devel-4.4.1_1", PkgVer{"libxmp-devel", "4.4.1", 1}},
		pkgVerSuccCase{"log4cplus-32bit-2.0.3_1", PkgVer{"log4cplus-32bit", "2.0.3", 1}},
		pkgVerSuccCase{"lsp-0.2.0.20160318_1", PkgVer{"lsp", "0.2.0.20160318", 1}},
		pkgVerSuccCase{"lxrandr-0.3.1_1", PkgVer{"lxrandr", "0.3.1", 1}},
		pkgVerSuccCase{"make-4.2.1_4", PkgVer{"make", "4.2.1", 4}},
		pkgVerSuccCase{"midori-7.0_1", PkgVer{"midori", "7.0", 1}},
		pkgVerSuccCase{"mimic-devel-32bit-1.2.0.2_3", PkgVer{"mimic-devel-32bit", "1.2.0.2", 3}},
		pkgVerSuccCase{"minizip-32bit-1.2.11_3", PkgVer{"minizip-32bit", "1.2.11", 3}},
		pkgVerSuccCase{"miro-video-converter-3.0.2_2", PkgVer{"miro-video-converter", "3.0.2", 2}},
		pkgVerSuccCase{"mit-krb5-devel-1.16.3_1", PkgVer{"mit-krb5-devel", "1.16.3", 1}},
		pkgVerSuccCase{"mupdf-tools-1.14.0_2", PkgVer{"mupdf-tools", "1.14.0", 2}},
		pkgVerSuccCase{"nyx-2.1.0_1", PkgVer{"nyx", "2.1.0", 1}},
		pkgVerSuccCase{"ofono-devel-1.28_1", PkgVer{"ofono-devel", "1.28", 1}},
		pkgVerSuccCase{"opendoas-6.0_1", PkgVer{"opendoas", "6.0", 1}},
		pkgVerSuccCase{"oprofile-1.3.0_1", PkgVer{"oprofile", "1.3.0", 1}},
		pkgVerSuccCase{"oracle-jdk-8u192_1", PkgVer{"oracle-jdk", "8u192", 1}},
		pkgVerSuccCase{"pango-view-1.42.4_2", PkgVer{"pango-view", "1.42.4", 2}},
		pkgVerSuccCase{"pax-20171021_2", PkgVer{"pax", "20171021", 2}},
		pkgVerSuccCase{"pcsc-ccid-1.4.30_2", PkgVer{"pcsc-ccid", "1.4.30", 2}},
		pkgVerSuccCase{"perl-HTTP-MultiPartParser-0.02_1", PkgVer{"perl-HTTP-MultiPartParser", "0.02", 1}},
		pkgVerSuccCase{"perl-Set-IntSpan-1.19_2", PkgVer{"perl-Set-IntSpan", "1.19", 2}},
		pkgVerSuccCase{"perl-Sub-Identify-0.14_4", PkgVer{"perl-Sub-Identify", "0.14", 4}},
		pkgVerSuccCase{"perl-Types-Serialiser-1.0_2", PkgVer{"perl-Types-Serialiser", "1.0", 2}},
		pkgVerSuccCase{"pipewire-0.2.5_1", PkgVer{"pipewire", "0.2.5", 1}},
		pkgVerSuccCase{"poppler-qt4-devel-32bit-0.61.1_1", PkgVer{"poppler-qt4-devel-32bit", "0.61.1", 1}},
		pkgVerSuccCase{"potrace-devel-32bit-1.15_1", PkgVer{"potrace-devel-32bit", "1.15", 1}},
		pkgVerSuccCase{"python-PyQt5-serialport-5.11.3_3", PkgVer{"python-PyQt5-serialport", "5.11.3", 3}},
		pkgVerSuccCase{"python3-ansicolor-0.2.6_2", PkgVer{"python3-ansicolor", "0.2.6", 2}},
		pkgVerSuccCase{"python3-pygame-1.9.4_2", PkgVer{"python3-pygame", "1.9.4", 2}},
		pkgVerSuccCase{"qt5-translations-5.11.3_4", PkgVer{"qt5-translations", "5.11.3", 4}},
		pkgVerSuccCase{"qt5-virtualkeyboard-32bit-5.11.3_4", PkgVer{"qt5-virtualkeyboard-32bit", "5.11.3", 4}},
		pkgVerSuccCase{"qutebrowser-1.5.2_1", PkgVer{"qutebrowser", "1.5.2", 1}},
		pkgVerSuccCase{"rcs-5.9.4_3", PkgVer{"rcs", "5.9.4", 3}},
		pkgVerSuccCase{"ruby-launchy-2.4.3_2", PkgVer{"ruby-launchy", "2.4.3", 2}},
		pkgVerSuccCase{"rxvt-unicode-9.22_9", PkgVer{"rxvt-unicode", "9.22", 9}},
		pkgVerSuccCase{"skb-0.4_1", PkgVer{"skb", "0.4", 1}},
		pkgVerSuccCase{"slurm-wlm-32bit-18.08.3.1_2", PkgVer{"slurm-wlm-32bit", "18.08.3.1", 2}},
		pkgVerSuccCase{"snooze-0.3_1", PkgVer{"snooze", "0.3", 1}},
		pkgVerSuccCase{"solarus-1.6.0_1", PkgVer{"solarus", "1.6.0", 1}},
		pkgVerSuccCase{"sound-theme-freedesktop-0.8_2", PkgVer{"sound-theme-freedesktop", "0.8", 2}},
		pkgVerSuccCase{"sqlcipher-devel-4.0.1_2", PkgVer{"sqlcipher-devel", "4.0.1", 2}},
		pkgVerSuccCase{"switchboard-plug-printers-32bit-2.1.6_1", PkgVer{"switchboard-plug-printers-32bit", "2.1.6", 1}},
		pkgVerSuccCase{"sword-1.8.1_3", PkgVer{"sword", "1.8.1", 3}},
		pkgVerSuccCase{"testdisk-7.0_4", PkgVer{"testdisk", "7.0", 4}},
		pkgVerSuccCase{"tig-2.4.1_1", PkgVer{"tig", "2.4.1", 1}},
		pkgVerSuccCase{"usbredir-0.8.0_1", PkgVer{"usbredir", "0.8.0", 1}},
		pkgVerSuccCase{"virtme-0.0.3.20180725_1", PkgVer{"virtme", "0.0.3.20180725", 1}},
		pkgVerSuccCase{"xfce4-wavelan-plugin-0.6.0_2", PkgVer{"xfce4-wavelan-plugin", "0.6.0", 2}},
		pkgVerSuccCase{"yelp-tools-3.28.0_1", PkgVer{"yelp-tools", "3.28.0", 1}},
		pkgVerSuccCase{"znc-32bit-1.7.2_1", PkgVer{"znc-32bit", "1.7.2", 1}},
		pkgVerSuccCase{"zziplib-devel-32bit-0.13.69_2", PkgVer{"zziplib-devel-32bit", "0.13.69", 2}},

		// Invalid strings
		pkgVerFailCase{"", ErrPkgVerNoRevision},
		pkgVerFailCase{"foobar-1", ErrPkgVerNoRevision},
		pkgVerFailCase{"foobar-1_0", ErrPkgVerBadRevision},
		pkgVerFailCase{"foobar-1_-5", ErrPkgVerBadRevision},
		pkgVerFailCase{"foobar-1_f", ErrPkgVerBadRevision},
		pkgVerFailCase{"foobar_1", ErrPkgVerNoVersion},
		pkgVerFailCase{"foobar-_1", ErrPkgVerNoVersion},
		pkgVerFailCase{"foobar-1:0_1", ErrPkgVerMalformedVersion},
		pkgVerFailCase{"foobar-1-_1", ErrPkgVerNoVersion},
		pkgVerFailCase{"-123_1", ErrPkgVerNoName},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name(), c.run)
	}
}
