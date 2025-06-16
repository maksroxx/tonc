# tonc — CLI для компиляции FunC контрактов

Утилита `tonc` предназначена для удобной компиляции и упаковки смарт-контрактов на языке FunC в формат BOC и JSON.  
Реализована на Go и использует официальные инструменты `func` и `fift` из TON SDK.

---

## Возможности

- Компиляция одного контракта или всех контрактов из папки
- Генерация `.cell.boc` бинарного файла
- Экспорт hex-представления BOC в JSON файл
- Удобный интерфейс с флагами и командной структурой
- Подробный лог и вывод прогресса

---

## Установка

1. Клонируйте репозиторий:

```bash
git clone https://github.com/maksroxx/tonc
cd tonc
```

2. Соберите бинарник:
```bash
make build
```

3. Убедитесь, что в системе установлены и доступны из PATH:
* func — компилятор FunC
* fift — Fift интерпретатор

4. Настройте переменную окружения FIFTPATH, чтобы fift находил свои библиотеки:
```bash
export FIFTPATH=/opt/homebrew/Cellar/ton/64/lib/fift
```

---

## Использование
```bash
./tonc build --contract ./contracts/contract.fc --boc --json --hex --verbose
```
### Команда build
* --contract — путь к одному .fc файлу для компиляции
* --src — путь к папке с .fc файлами для пакетной компиляции
* --boc — сгенерировать .cell.boc файл
* --json — сгенерировать .compiled.json файл с hex-представлением BOC
* --hex — включить вывод hex-строки в JSON
* --verbose — подробный вывод

---

## Примеры
* Компиляция одного контракта:
```bash
./tonc build --contract ./contracts/contract.fc --boc --json --hex
```
* Компиляция всех контрактов из папки contracts:
```bash
./tonc build --src ./contracts --boc --json --hex
```

---

## Makefile
Доступны команды:
* make build — собрать бинарник CLI
* make run — собрать и запустить CLI на примере контракта
* make clean — удалить артефакты сборки и бинарник
* make compile-all — скомпилировать все контракты из папки contracts

---

## Требования
* Go 1.18+
* Установленный TON SDK с func и fift
* Правильно настроенная переменная окружения FIFTPATH