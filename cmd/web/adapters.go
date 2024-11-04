package main

//import (
//	"encoding/json"
//	"encoding/xml"
//	"go-breeders/models"
//	"io"
//	"log"
//	"net/http"
//)
//
//type CatBreedsInterface interface {
//	GetAllCatBreeds() ([]*models.CatBreed, error)
//}
//
//type RemoteService struct {
//	Remote CatBreedsInterface
//}
//
//func (rs *RemoteService) GetAllBreeds() ([]*models.CatBreed, error) {
//	return rs.Remote.GetAllCatBreeds()
//}
//
//type JSONBackend struct {
//}
//
//type XMLBackend struct {
//}
//
//func (jb *JSONBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
//	var breeds []*models.CatBreed
//	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/json")
//	if err != nil {
//		return nil, err
//	}
//
//	defer func(Body io.ReadCloser) {
//		err = Body.Close()
//		if err != nil {
//			log.Println(err)
//		}
//	}(resp.Body)
//
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	err = json.Unmarshal(body, &breeds)
//	if err != nil {
//		return nil, err
//	}
//	return breeds, nil
//}
//
//func (xb *XMLBackend) GetAllCatBreeds() ([]*models.CatBreed, error) {
//
//	resp, err := http.Get("http://localhost:8081/api/cat-breeds/all/xml")
//	if err != nil {
//		return nil, err
//	}
//	defer resp.Body.Close()
//	body, err := io.ReadAll(resp.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	type catBreeds struct {
//		XMLName xml.Name           `xml:"cat-breeds"`
//		Breeds  []*models.CatBreed `xml:"cat-breed"`
//	}
//
//	var breeds catBreeds
//
//	err = xml.Unmarshal(body, &breeds)
//	if err != nil {
//		return nil, err
//	}
//	return breeds.Breeds, nil
//}
