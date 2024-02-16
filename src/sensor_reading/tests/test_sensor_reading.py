from sensor_reading.sensor_reading import read_sensor_data_from_csv, generate_data, connect_mqtt  # Ajuste os imports conforme necessário

def test_read_sensor_data_from_csv():
    generator = read_sensor_data_from_csv("sensors.csv")
    first_row = next(generator)
    assert 'radiacao_solar' in first_row and 'temperatura' in first_row, "As chaves esperadas não estão presentes no CSV"

def test_generate_data():
    radiacao_solar = generate_data("radiacao_solar")
    assert radiacao_solar is not None, "A função generate_data deve retornar um valor que não seja None"
