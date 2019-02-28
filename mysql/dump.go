// Package mysql ...
package mysql

//	"github.com/lucashtc/gobackup/execmysql"

//"github.com/lucashtc/gobackup/mysql"

// DirDump function create directory for save files backup
// Separando por pastas cadas tipo de objeto do schema
/*func DirDump(local, time, dataBase string) (string, error) {
	d, err := dir.CreateDir(time, dataBase, local)
	if err != nil {
		return "", err
	}
	return d, nil
}

// DumpAll function vai realizar dump de toda o Schema encontrado no servidor
func DumpAll(cf *DataBase) {

	time := helper.GetCurrentTime()

	Log(cf.Dir, "", LIMIT)
	Log(cf.Dir, "", LIMIT)

	Log(cf.Dir, helper.GetCurrentTime(), SEARCHSCHEMA)
	db, err := GetData(cf)
	if err != nil {
		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Falha ao obter dados da bases >>>> \n %s", err))
	}

	Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Criando dumps... \n\n"))
	for _, d := range db {
		var w sync.WaitGroup
		if d.Name == "" {
			//Caso o Name seja empty pula para a proxima execução do For
			continue
		}
		cf.Name = d.Name
		Log(cf.Dir, helper.GetCurrentTime(), LIMIT)

		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Fazendo backup da base %s\n", d.Name))
		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Criando pasta da base %s\n", d.Name))

		dirName, err := DirDump(cf.Dir, time, d.Name)
		if err != nil {
			Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Falha ao criar pasta %s \n error: >>>> error %s \n", d.Name, err))
			continue
		}

		cf.DirFile = fmt.Sprintf("%s/%s.sql", dirName, d.Name)

		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Realizando dump da base %s \n", d.Name))

		// Dump for database
		if err := dumpdb(cf); err != nil {
			Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprint(err))
			continue
		}

		Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Compactando arquivo da base \n"))

		w.Add(1)
		go func() {
			if err := helper.Zip(cf.DirFile); err != nil {
				Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf(ZIPFAIL, cf.DirFile))
			}
			w.Done()
		}()
		w.Wait()

		if err := dir.Delete(cf.DirFile); err != nil {
			Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf("Error: %s", err))
			continue
		}
		Log(cf.Dir, helper.GetCurrentTime(), LIMIT)
	}

	Log(cf.Dir, helper.GetCurrentTime(), fmt.Sprintf(" Fim da execução do script de backup\n\n"))
}

// dumpdb Dump database
func dumpdb(conf *DataBase) error {
	var param []string
	param = append(param, "-u", conf.User)
	if conf.Password != "" {
		pass := fmt.Sprintf("-p%s", conf.Password)
		param = append(param, pass)
	}
	param = append(param, "--skip-add-drop-table")
	param = append(param, conf.Name)
	param = append(param, "-r", conf.DirFile)

	_, err := execmysql.ExecDump(param)
	if err != nil {
		return fmt.Errorf("Falha ao executar dump da base %s \n Error >> %s", conf.Name, err)
	}
	return nil
}

// Log ...
func Log(name, time, text string) {
	fmt.Print(text)

	if err := os.MkdirAll(name, os.ModePerm); err != nil {
		fmt.Println(err)
	}
	local := fmt.Sprintf("%s/log.txt", name)
	if err := dir.CreateFile(local); err != nil {
		fmt.Printf("Falha ao criar arquivo. Error %s\n", err)
	}

	text = fmt.Sprintf("%s========%s\r\n", time, text)
	if err := dir.WriteFile(local, text); err != nil {
		fmt.Printf("Falha ao escrever no arquivo %s. Error %s\n", local, err)
	}
}

var (
	// LIMIT ...
	LIMIT = "=================================================== \n"

	// ZIPFAIL ...
	ZIPFAIL = "Ocorreu um error ao compactar o arquivo %s \n"

	// SEARCHSCHEMA ...
	SEARCHSCHEMA = "Buscando informações dos schemas no servidor ...\n"
)
*/
