import 'package:dio/dio.dart';
import 'note.dart';

class ApiService {
  final Dio _dio = Dio();

  // Получение всех квартир
  Future<List<Note>> getApartments() async {
    try {
      final response = await _dio.get('http://192.168.1.78:8080/apartments');
      if (response.statusCode == 200) {
        return (response.data as List)
            .map((apartment) => Note.fromJson(apartment))
            .toList();
      } else {
        throw Exception('Failed to load apartments');
      }
    } catch (e) {
      throw Exception('Error fetching apartments: $e');
    }
  }
  Future<void> updateApartment(Note note) async {
    final data = {
      "ID": note.id, // Заглавные буквы
      "Title": note.title,
      "Description": note.description,
      "ImageLink": note.photo_id, // Заглавные буквы
      "Price": note.price,
    };

    print("Sending PUT request to update apartment...");
    print("URL: 'http://192.168.1.78:8080/apartments/update/${note.id}");
    print("Data: $data");

    try {
      final response = await _dio.put(
        'http://192.168.1.78:8080/apartments/update/${note.id}',
        data: data,
      );

      if (response.statusCode != 200) {
        throw Exception('Failed to update apartment');
      }
    } catch (e) {
      print('Error updating apartment: $e');
      throw Exception('Error updating apartment: $e');
    }
  }




  // Переключение избранного
  Future<void> toggleFavourite(int id) async {
    try {
      await _dio.put('http://192.168.1.78:8080/apartments/favourite/$id');
    } catch (e) {
      throw Exception('Error toggling favourite: $e');
    }
  }
}