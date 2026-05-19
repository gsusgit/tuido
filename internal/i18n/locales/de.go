package locales

import "github.com/gsus/todo-app/tui/internal/i18n"

func init() {
	i18n.Register("de", map[i18n.Key]string{
		i18n.KeyNewTask: "Neu", i18n.KeyFilter: "Filtern", i18n.KeyFilterActive: "Filter",
		i18n.KeyEdit: "Bearbeiten", i18n.KeyDelete: "Löschen", i18n.KeyComplete: "Erledigen",
		i18n.KeyNavigate: "Navigieren", i18n.KeyTheme: "Thema", i18n.KeyExit: "Beenden",
		i18n.KeySearch: "Suchen", i18n.KeyEmptyList: "Keine Aufgaben. Drücke 'n'.",
		i18n.KeyEmptyFiltered: "Keine passenden Aufgaben.",
		i18n.KeyTitleRequired: "Titel darf nicht leer sein",
		i18n.KeyCompletedOf: "Erledigt: %d von %d", i18n.KeyScrollMore: "Mehr anzeigen (scrollen oder %s)",
		i18n.KeyTitle: "Titel", i18n.KeyCategory: "Kategorie", i18n.KeyPriority: "Priorität",
		i18n.KeyStatus: "Status", i18n.KeyPending: "Offen", i18n.KeyStatusCompleted: "Erledigt",
		i18n.KeyFilterTitle: "Aufgaben filtern", i18n.KeyAll: "Alle",
		i18n.KeyTabNext: "Weiter", i18n.KeyEnterSave: "Speichern", i18n.KeyEnterApply: "Anwenden",
		i18n.KeyEscCancel: "Abbrechen", i18n.KeyBack: "Zurück", i18n.KeyReset: "Zurücksetzen",
		i18n.KeyTerminalSmall:       "Terminal zu klein — vergrößern",
		i18n.KeyWindowTooSmall:      "Fenster zu klein",
		i18n.KeyMinimumSizeRequired: "Mindestgröße: %d×%d",
		i18n.KeyDeletePending: "Löschen? (Enter bestätigen · Esc abbrechen)",
		i18n.KeyTaskCreated: "Aufgabe erstellt", i18n.KeyTaskSaved: "Gespeichert",
		i18n.KeyTaskDeleted: "Gelöscht", i18n.KeySaveError: "Fehler: %s",
		i18n.KeyLangChanged: "Sprache: %s", i18n.KeyResetPrompt: "Alle Daten in %s löschen? (j/N): ",
		i18n.KeyResetDone: "Daten gelöscht.", i18n.KeyResetCancelled: "Abgebrochen.",
		i18n.KeyResetInvalid: "Unbekannte Sprache: es, en, fr, de, it, pt",
		i18n.KeyPersonal: "Persönlich", i18n.KeyTrabajo: "Arbeit",
		i18n.KeyPriorityHigh: "Hoch", i18n.KeyPriorityMedium: "Mittel", i18n.KeyPriorityLow: "Niedrig",
		i18n.KeySelectLang: "Sprache wählen:", i18n.KeyLangNav: "↑↓ · Enter · Esc",
		i18n.KeyControls: "Steuerung", i18n.KeyHelpMenu: "Steuerung anzeigen",
		i18n.KeySort: "Sortierung", i18n.KeySortField: "Sortieren nach", i18n.KeySortDirection: "Richtung",
		i18n.KeySortDefault: "Standard", i18n.KeySortAsc: "Aufsteigend", i18n.KeySortDesc: "Absteigend",
	})
}
