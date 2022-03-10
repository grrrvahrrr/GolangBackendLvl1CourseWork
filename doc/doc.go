//The app is an URL shortener - it makes URLs shorter.
//
//Server can be configured using config.env file found in cmd/config
//
//It implements hexagonal architecture.
//On the inner layer there are a couple of business functions:
//Generates random short strings
//  func GenerateRandomString() string {
//  	var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
//  	b := make([]rune, 10)
//		for i := range b {
// 	 		b[i] = symbols[rand.Intn(len(symbols))]
// 		}
// 		return string(b)
//  }
//
//UpdateNumOfUses
//  func UpdateNumOfUses(data string) (string, error) {
// 		iv, err := strconv.Atoi(data)
// 		if err != nil {
// 			return "", err
// 		}
// 		iv++
// 		data = strconv.Itoa(iv)
//	 	return data, nil
//  }
//
//Package dbbackend connects storage and frontend while implementing businness logic.
//The package has following interface:
//  type DataStore interface {
// 		WriteURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error)
// 		WriteData(ctx context.Context, url entities.UrlData) (*entities.UrlData, error)
// 		ReadURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error)
// 		GetIPData(ctx context.Context, url entities.UrlData) (string, error)
//  }
//
//The storage itself is based on goleveldb and keeps all the information in local files.
//Function example:
//  func (fd *FullDataFile) WriteData(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {
// 		err := fd.datadb.Put([]byte(url.ShortURL), []byte(url.Data), nil)
// 		if err != nil {
// 			return nil, fmt.Errorf("error writing to datadb : %w", err)
// 		}
// 		d := strings.Join([]string{fd.URLData.ShortURL, fd.URLData.IP}, ":")
// 		err = fd.ipdb.Put([]byte(d), []byte(url.IPData), nil)
// 		if err != nil {
// 			return nil, fmt.Errorf("error writing to ipdb : %w", err)
// 		}
// 		return &url, nil
//  }
//
//The router part is generated from Open API file and is based on go-chi router.
//
//Redirect function:
//(GET /su/{shortURL})
//  func (rt *OpenApiChi) Redirect(w http.ResponseWriter, r *http.Request, shortURL string) {
// 		if shortURL == "" {
// 			err := render.Render(w, r, apichi.ErrInvalidRequest(http.ErrNotSupported))
// 			log.Error(err)
// 			if err != nil {
// 				log.Error(err)
// 			}
// 			return
// 		}
// 		ip, _, err := net.SplitHostPort(r.RemoteAddr)
// 		if err != nil {
// 			log.Error(err)
// 		}
// 		nud, err := rt.hs.RedirectionHandle(r.Context(), shortURL, ip)
// 		if err != nil {
// 			log.Error(errors.Unwrap(err))
// 			err = render.Render(w, r, apichi.ErrRender(errors.Unwrap(err)))
// 			if err != nil {
// 				log.Error(err)
// 			}
// 			return
// 		}
// 		http.Redirect(w, r, nud.FullURL, http.StatusSeeOther)
//  }
package doc
