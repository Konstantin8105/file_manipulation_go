package fileManipulation

import "strings"

var extentions = []string{
	".bmd", ".CFR", ".cod",
	".cut", ".day", ".dbi",
	".dbs", ".dgn", ".dsp",
	".slg", ".ecf", ".ejt",
	".EQL", ".est", ".EU2",
	".emf", ".NLD", ".num",
	".rea", ".REI_SPRO_Auxilary_Data",
	".sbk", ".scn", ".slv",
	".ANL", ".u01", ".u02",
	".u03", ".u04", ".u05",
	".u06", ".u07", ".u08",
	".UID", ".REI_Saved_Picture",
	".err", ".bsh", ".cfc",
	".ben", ".MMTX", ".dsn",
	".jst", ".REI_STEELDESIGNER_DATA",
	".str", "._MASS.TXT", ".ed.Backup",
	".msh", ".log",
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
