class Note {
  final int id;
  final String title;
  final String description;
  final String photo_id; // Поле для фото
  final double price;
  bool isFavorite;

  Note({
    required this.id,
    required this.title,
    required this.description,
    required this.photo_id, // Исправлено
    required this.price,
    this.isFavorite = false,
  });


  // Обработка JSON-ответа
  factory Note.fromJson(Map<String, dynamic> json) {
    return Note(
      id: json['ID'] ?? 0, // Если null, заменить на 0
      title: json['Title'] ?? 'Нет названия', // Если Title null
      description: json['Description'] ?? 'Описание отсутствует', // Если Description null
      photo_id: json['ImageLink'] ?? '', // Если ImageLink null, заменить на пустую строку
      price: (json['Price'] ?? 0).toDouble(), // Если Price null, заменить на 0.0
      isFavorite: json['Favourite'] ?? false, // Если Favourite null, заменить на false
    );
  }
}