import 'package:flutter/material.dart';
import '../models/note.dart';
import 'edit_note_page.dart';

class NotePage extends StatelessWidget {
  final Note note;
  final Function(Note) onEdit;

  const NotePage({
    Key? key,
    required this.note,
    required this.onEdit,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(note.title),
        actions: [
          IconButton(
            icon: const Icon(Icons.edit),
            onPressed: () {
              // Переход на экран редактирования
              Navigator.of(context).push(
                MaterialPageRoute(
                  builder: (context) => EditNotePage(
                    note: note,
                    onSave: (updatedNote) {
                      onEdit(updatedNote); // Обновляем родительский виджет
                      Navigator.of(context).pop(); // Закрываем экран редактирования
                    },
                  ),
                ),
              );
            },
          ),
        ],
      ),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            if (note.photo_id.isNotEmpty)
              Image.network(
                note.photo_id,
                fit: BoxFit.cover,
                errorBuilder: (context, error, stackTrace) {
                  return const Text('Ошибка загрузки изображения');
                },
              ),
            const SizedBox(height: 16.0),
            Text(
              'Цена: ₽${note.price.toStringAsFixed(2)}',
              style: const TextStyle(fontSize: 20, fontWeight: FontWeight.bold),
            ),
            const SizedBox(height: 16.0),
            Text(
              note.description,
              style: const TextStyle(fontSize: 16),
            ),
          ],
        ),
      ),
    );
  }
}