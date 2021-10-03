# golang_gui_easy_widgets
# НАХОДИТСЯ В РАЗРАБОТКЕ !

На данный момент вы можете реализовать простейшие виджеты и далее взаимодействовать с ними
Разместите как простейшие кнопки так и кастомные со своим отбражением,размещайте как текст, так
и картинки, визуализируйте свой прогресс с помощью ProgressBar, выводите графики функций и 
другие данные на координатную плоскость CoordinateSurface, используйте столбчатые диаграммы и поля для ввода
текста - это добавит динамичности и работу с клиентом. Дополнительно имеется структура FractalMath
для получения координат и отрисовки точками на поверхности треугольник Серпинского.

Пример использования:

# -------------------------------------------------------------------

package main

import (
	"github.com/gonutz/prototype/draw"
	"github.com/toolsgui/tools2D"
)

// имя переменной = виджет(парамеры)
var NAMEVAR = tools2D.[WIDGET](..PARAMS)

// вункция обновления
func FUNCTIONUPDATE(win draw.Window) {
	NAMEVAR.View(win)
}

// точка входа
func main() {
  // создание окна (имя заголовка, ширина, высота, функция обновления)
	draw.RunWindow("Name window", WidthSize, HeightsSise, FUNCTIONUPDATE)
} 

# -------------------------------------------------------------------

Вам нужно:

--> Создать виджеты и глобально сохранить в переменные.
--> Создать точку входа (функцию Main) 
--> Создать функцию для обновления окна и всего что в нем расположите
