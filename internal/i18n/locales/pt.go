package locales

import "github.com/gsusgit/tuido/internal/i18n"

func init() {
	i18n.Register("pt", map[i18n.Key]string{
		i18n.KeyNewTask: "Nova", i18n.KeyFilter: "Filtrar", i18n.KeyFilterActive: "Filtro",
		i18n.KeyEdit: "Editar", i18n.KeyDelete: "Eliminar", i18n.KeyComplete: "Concluir",
		i18n.KeyNavigate: "Navegar", i18n.KeyTheme: "Escolher tema", i18n.KeyThemeTitle: "Escolher tema", i18n.KeyExit: "Sair",
		i18n.KeySearch: "Pesquisar", i18n.KeyEmptyList: "Sem tarefas. Pressione 'n'.",
		i18n.KeyEmptyFiltered: "Nenhuma tarefa correspondente.",
		i18n.KeyTitleRequired: "O título não pode estar vazio",
		i18n.KeyCompletedOf: "Concluídas: %d de %d", i18n.KeyScrollMore: "Ver mais (scroll ou %s)",
		i18n.KeyTitle: "Título", i18n.KeyCategory: "Categoria", i18n.KeyPriority: "Prioridade",
		i18n.KeyStatus: "Estado", i18n.KeyPending: "Pendentes", i18n.KeyStatusCompleted: "Concluídas",
		i18n.KeyFilterTitle: "Filtrar tarefas", i18n.KeyAll: "Todas",
		i18n.KeyTabNext: "Seguinte", i18n.KeyEnterSave: "Guardar", i18n.KeyEnterApply: "Aplicar",
		i18n.KeyEscCancel: "Cancelar", i18n.KeyBack: "Voltar", i18n.KeyReset: "Repor",
		i18n.KeyTerminalSmall:       "Terminal demasiado pequeno — redimensione",
		i18n.KeyWindowTooSmall:      "Janela demasiado pequena",
		i18n.KeyMinimumSizeRequired: "Tamanho mínimo: %d×%d",
		i18n.KeyDeletePending: "Eliminar? (Enter confirmar · Esc cancelar)",
		i18n.KeyTaskCreated: "Tarefa criada", i18n.KeyTaskSaved: "Guardada",
		i18n.KeyTaskDeleted: "Eliminada", i18n.KeySaveError: "Erro: %s",
		i18n.KeyLangChanged: "Idioma: %s", i18n.KeyResetPrompt: "Apagar dados em %s? (s/N): ",
		i18n.KeyResetDone: "Dados apagados.", i18n.KeyResetCancelled: "Cancelado.",
		i18n.KeyResetInvalid: "Idioma desconhecido: es, en, fr, de, it, pt",
		i18n.KeyPersonal: "Pessoal", i18n.KeyTrabajo: "Trabalho",
		i18n.KeyPriorityHigh: "Alta", i18n.KeyPriorityMedium: "Média", i18n.KeyPriorityLow: "Baixa",
		i18n.KeySelectLang: "Selecionar idioma:", i18n.KeyLangNav: "↑↓ · Enter · Esc",
		i18n.KeyControls: "Controlos", i18n.KeyHelpMenu: "Ver controlos",
		i18n.KeySort: "Ordenação", i18n.KeySortField: "Ordenar por", i18n.KeySortDirection: "Direção",
		i18n.KeySortDefault: "Predefinido", i18n.KeySortAsc: "Ascendente", i18n.KeySortDesc: "Descendente",
	})
}
