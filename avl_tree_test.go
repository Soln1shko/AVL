package main

import (
	"testing"
)

// Вспомогательная функция для проверки баланса дерева
func isBalanced(node *Node) bool {
	if node == nil {
		return true
	}

	balance := node.BalanceFactor()
	if balance < -1 || balance > 1 {
		return false
	}

	return isBalanced(node.Left) && isBalanced(node.Right)
}

// Вспомогательная функция для проверки свойства BST
func isBST(node *Node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Key <= min || node.Key >= max {
		return false
	}

	return isBST(node.Left, min, node.Key) && isBST(node.Right, node.Key, max)
}

func TestAVLInsert(t *testing.T) {
	tree := &AVLTree{}

	// Тест 1: Вставка в пустое дерево
	tree.Insert(10)
	if tree.Root == nil || tree.Root.Key != 10 {
		t.Error("Ошибка при вставке в пустое дерево")
	}

	// Тест 2: Простая вставка без необходимости балансировки
	tree.Insert(20)
	if tree.Root.Right == nil || tree.Root.Right.Key != 20 {
		t.Error("Ошибка при простой вставке")
	}

	// Тест 3: Вставка, требующая правого поворота
	tree.Insert(5)
	tree.Insert(3)
	if !isBalanced(tree.Root) {
		t.Error("Дерево не сбалансировано после вставки")
	}

	// Тест 4: Проверка свойства BST
	if !isBST(tree.Root, -1000000, 1000000) {
		t.Error("Нарушено свойство бинарного дерева поиска")
	}
}

func TestAVLRotations(t *testing.T) {
	tree := &AVLTree{}

	// Тест левого поворота
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)

	if tree.Root.Key != 20 {
		t.Error("Ошибка при левом повороте")
	}

	if !isBalanced(tree.Root) {
		t.Error("Дерево не сбалансировано после левого поворота")
	}

	// Тест право-левого поворота
	tree = &AVLTree{} // Новое дерево
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)

	if tree.Root.Key != 15 {
		t.Error("Ошибка при право-левом повороте")
	}

	if !isBalanced(tree.Root) {
		t.Error("Дерево не сбалансировано после право-левого поворота")
	}
}

func TestAVLHeight(t *testing.T) {
	tree := &AVLTree{}

	// Проверка высоты пустого дерева
	if tree.Root.GetHeight() != 0 {
		t.Error("Неверная высота пустого дерева")
	}

	// Вставка элементов и проверка высоты
	tree.Insert(10)
	if tree.Root.GetHeight() != 1 {
		t.Error("Неверная высота дерева с одним узлом")
	}

	tree.Insert(20)
	tree.Insert(30)

	// После балансировки высота должна быть 2
	if tree.Root.GetHeight() != 2 {
		t.Error("Неверная высота после балансировки")
	}
}

func TestAVLBalanceFactor(t *testing.T) {
	tree := &AVLTree{}

	// Проверка фактора баланса для разных конфигураций
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)

	if tree.Root.BalanceFactor() != 0 {
		t.Error("Неверный фактор баланса для сбалансированного дерева")
	}

	tree.Insert(3)
	if abs(tree.Root.BalanceFactor()) > 1 {
		t.Error("Превышен допустимый фактор баланса")
	}
}

// Вспомогательная функция для вычисления модуля числа
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
