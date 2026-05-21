package locales

import "github.com/gsusgit/tuido/internal/i18n"

func init() {
	i18n.Register("it", map[i18n.Key]string{
		i18n.KeyNewTask: "Nuova", i18n.KeyFilter: "Filtra", i18n.KeyFilterActive: "Filtro",
		i18n.KeyEdit: "Modifica", i18n.KeyDelete: "Elimina", i18n.KeyComplete: "Completa",
		i18n.KeyNavigate: "Naviga", i18n.KeyTheme: "Scegli tema", i18n.KeyThemeTitle: "Scegli un tema", i18n.KeyExit: "Esci",
		i18n.KeySearch: "Cerca", i18n.KeyEmptyList: "Nessuna attività. Premi 'n'.",
		i18n.KeyEmptyFiltered: "Nessuna corrispondenza.",
		i18n.KeyTitleRequired: "Il titolo non può essere vuoto",
		i18n.KeyCompletedOf: "Completate: %d di %d", i18n.KeyScrollMore: "Mostra altro (scroll o %s)",
		i18n.KeyTitle: "Titolo", i18n.KeyCategory: "Categoria", i18n.KeyPriority: "Priorità",
		i18n.KeyStatus: "Stato", i18n.KeyPending: "In sospeso", i18n.KeyStatusCompleted: "Completate",
		i18n.KeyFilterTitle: "Filtra attività", i18n.KeyAll: "Tutte",
		i18n.KeyTabNext: "Avanti", i18n.KeyEnterSave: "Salva", i18n.KeyEnterApply: "Applica",
		i18n.KeyEscCancel: "Annulla", i18n.KeyBack: "Indietro", i18n.KeyReset: "Reset",
		i18n.KeyTerminalSmall:       "Terminale troppo piccolo — ridimensiona",
		i18n.KeyWindowTooSmall:      "Finestra troppo piccola",
		i18n.KeyMinimumSizeRequired: "Dimensione minima: %d×%d",
		i18n.KeyDeletePending: "Eliminare? (Invio conferma · Esc annulla)",
		i18n.KeyTaskCreated: "Attività creata", i18n.KeyTaskSaved: "Salvata",
		i18n.KeyTaskDeleted: "Eliminata", i18n.KeySaveError: "Errore: %s",
		i18n.KeyLangChanged: "Lingua: %s", i18n.KeyResetPrompt: "Eliminare %s? (s/N): ",
		i18n.KeyResetDone: "Dati eliminati.", i18n.KeyResetCancelled: "Annullato.",
		i18n.KeyResetInvalid: "Lingua sconosciuta: es, en, fr, de, it, pt",
		i18n.KeyPersonal: "Personale", i18n.KeyTrabajo: "Lavoro",
		i18n.KeyPriorityHigh: "Alta", i18n.KeyPriorityMedium: "Media", i18n.KeyPriorityLow: "Bassa",
		i18n.KeySelectLang: "Seleziona lingua:", i18n.KeyLangNav: "↑↓ · Invio · Esc",
		i18n.KeyControls: "Controlli", i18n.KeyHelpMenu: "Mostra controlli",
		i18n.KeySort: "Ordine", i18n.KeySortField: "Ordina per", i18n.KeySortDirection: "Direzione",
		i18n.KeySortDefault: "Predefinito", i18n.KeySortAsc: "Crescente", i18n.KeySortDesc: "Decrescente",
	})
}
