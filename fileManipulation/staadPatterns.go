package fileManipulation

import "strings"

var extentions = []string{
	".ANL",
	".Accel_TIMEH.DAT",
	".Accel.mhs",
	".CFR",
	".EQL",
	".EU2",
	".LR2",
	".MMTX",
	".NLD",
	".REI_SPRO_Auxilary_Data",
	".REI_STEELDESIGNER_DATA",
	".REI_Saved_Picture",
	".Rei_Concrete_Link",
	".UID",
	".adfx",
	".anl",
	".ben",
	".bmd",
	".bsh",
	".cfc",
	".cod",
	".cut",
	".day",
	".dbi",
	".dbs",
	".dgn",
	".dsn",
	".dsp",
	".ecf",
	".ed.Backup",
	".ejt",
	".emf",
	".err",
	".ess",
	".est",
	".jst",
	".log",
	".log",
	".mhs",
	".njd",
	".num",
	".rea",
	".rsd",
	".sbk",
	".scn",
	".slg",
	".slv",
	".SLV",
	".std.metadata",
	".std.un~",
	".std~",
	".str",
	".u01",
	".u02",
	".u03",
	".u04",
	".u05",
	".u06",
	".u07",
	".u08",
	"_L19.TXT",
	"_L19.txt",
	"_SEC.L33",
	"_TIMEH.DAT",
	"_MASS.TXT",
	"_NONL_JOINT_DISP.TXT",
}

var staadFile = []string{
	".std",
}

var ignoreFolderName = []string{
	".svn", ".git",
}

func isStaadFile(filename string) bool {
	for _, ext := range staadFile {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

func isStaadTempFile(filename string) bool {
	for _, ext := range extentions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

func isIgnoreFolder(folderName string) bool {
	for _, ignore := range ignoreFolderName {
		if folderName == ignore {
			return true
		}
	}
	return false
}
