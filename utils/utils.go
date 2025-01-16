package utils

import (
	"context"
	"fmt"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

const GOOGLE_APPLICATION_CREDENTIALS = "GOOGLE_APPLICATION_CREDENTIALS" //AGREGAR
const ID_HOJA_DE_CALCULO = "ID_HOJA_DE_CALCULO"                         //AGREGAR

type Serve struct {
	*sheets.Service
}

func (s *Serve) ConfigCredential() error {
	ctx := context.Background()
	var err error
	s.Service, err = sheets.NewService(ctx, option.WithCredentialsFile(GOOGLE_APPLICATION_CREDENTIALS))
	if err != nil {
		return fmt.Errorf("No se pudo crear el cliente de sheets: %v", err)
	}
	return nil
}

func (s *Serve) AddGasto(servicio string, monto string) error {
	appendRange := "Hoja 1!A:A" // Usar A:A para agregar datos al final
	values := [][]interface{}{
		{servicio, monto},
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := s.Service.Spreadsheets.Values.Append(ID_HOJA_DE_CALCULO, appendRange, valueRange).ValueInputOption("RAW").InsertDataOption("INSERT_ROWS").Do()
	return err
}

func (s *Serve) AddService(servicio string) error {
	appendRange := "Hoja 1!A:A" // Usar A:A para agregar datos al final
	values := [][]interface{}{
		{servicio},
	}

	valueRange := &sheets.ValueRange{
		Values: values,
	}

	_, err := s.Service.Spreadsheets.Values.Append(ID_HOJA_DE_CALCULO, appendRange, valueRange).ValueInputOption("RAW").Do()
	return err
}

func (s *Serve) GetGastos() ([][]interface{}, error) {
	readRange := "Hoja 1!A1:Z1000" // Asegúrate de que el rango sea correcto

	resp, err := s.Service.Spreadsheets.Values.Get(ID_HOJA_DE_CALCULO, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("No se pudo leer la hoja de cálculo: %v", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("No hay datos en la hoja de cálculo")
	}

	return resp.Values, nil
}

func (s *Serve) GetService() ([]interface{}, error) {
	readRange := "Hoja 1!A:A" // Solo la columna A

	resp, err := s.Service.Spreadsheets.Values.Get(ID_HOJA_DE_CALCULO, readRange).Do()
	if err != nil {
		return nil, fmt.Errorf("No se pudo leer la hoja de cálculo: %v", err)
	}

	if len(resp.Values) == 0 {
		return nil, fmt.Errorf("No hay datos en la hoja de cálculo")
	}

	// Extraemos la primera columna
	var columnData []interface{}
	for _, row := range resp.Values {
		if len(row) > 0 {
			columnData = append(columnData, row[0]) // solo tomamos el primer valor de cada fila (columna A)
		}
	}

	return columnData, nil
}
