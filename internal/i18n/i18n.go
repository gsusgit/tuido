package i18n

import "fmt"

// Key identifies a translatable string.
type Key string

// Lang describes a supported language.
type Lang struct {
	Code string
	Name string
}

const (
	KeyNewTask        Key = "new_task"
	KeyFilter         Key = "filter"
	KeyFilterActive   Key = "filter_active"
	KeyEdit           Key = "edit"
	KeyDelete         Key = "delete"
	KeyComplete       Key = "complete"
	KeyNavigate       Key = "navigate"
	KeyTheme          Key = "theme"
	KeyExit           Key = "exit"
	KeySearch         Key = "search"
	KeyEmptyList      Key = "empty_list"
	KeyEmptyFiltered  Key = "empty_filtered"
	KeyTitleRequired  Key = "title_required"
	KeyCompletedOf    Key = "completed_of"
	KeyScrollMore     Key = "scroll_more"
	KeyTitle          Key = "title"
	KeyCategory       Key = "category"
	KeyPriority       Key = "priority"
	KeyStatus         Key = "status"
	KeyPending        Key = "pending"
	KeyStatusCompleted Key = "status_completed"
	KeyFilterTitle    Key = "filter_title"
	KeyAll            Key = "all"
	KeyTabNext        Key = "tab_next"
	KeyEnterSave      Key = "enter_save"
	KeyEnterApply     Key = "enter_apply"
	KeyEscCancel      Key = "esc_cancel"
	KeyBack           Key = "back"
	KeyReset          Key = "reset"
	KeyTerminalSmall         Key = "terminal_small"
	KeyWindowTooSmall      Key = "window_too_small"
	KeyMinimumSizeRequired Key = "minimum_size_required"
	KeyDeleteConfirm  Key = "delete_confirm"
	KeyDeletePending  Key = "delete_pending"
	KeyTaskCreated    Key = "task_created"
	KeyTaskSaved      Key = "task_saved"
	KeyTaskDeleted    Key = "task_deleted"
	KeySaveError      Key = "save_error"
	KeyLangChanged    Key = "lang_changed"
	KeyResetPrompt    Key = "reset_prompt"
	KeyResetDone      Key = "reset_done"
	KeyResetCancelled Key = "reset_cancelled"
	KeyResetInvalid   Key = "reset_invalid_lang"
	KeyHelpUsage      Key = "help_usage"
	KeyMigrating      Key = "migrating"
	KeyPersonal       Key = "personal"
	KeyTrabajo        Key = "trabajo"
	KeyPriorityHigh   Key = "priority_high"
	KeyPriorityMedium Key = "priority_medium"
	KeyPriorityLow    Key = "priority_low"
	KeySelectLang     Key = "select_lang"
	KeyLangNav        Key = "lang_nav"
	KeyControls       Key = "controls"
	KeyHelpMenu       Key = "help_menu"
	KeySort           Key = "sort"
	KeySortField      Key = "sort_field"
	KeySortDirection  Key = "sort_direction"
	KeySortDefault    Key = "sort_default"
	KeySortAsc        Key = "sort_asc"
	KeySortDesc       Key = "sort_desc"
)

var (
	current  = "en"
	catalogs = map[string]map[Key]string{}
)

// Register adds a language catalog.
func Register(lang string, catalog map[Key]string) {
	catalogs[lang] = catalog
}

// SetLang sets the active language if supported.
func SetLang(lang string) {
	if _, ok := catalogs[lang]; ok {
		current = lang
	}
}

// Current returns the active language code.
func Current() string {
	return current
}

// T returns the translated string for key, with optional fmt args.
func T(key Key, args ...any) string {
	msg := lookup(current, key)
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}

func lookup(lang string, key Key) string {
	if msg, ok := catalogs[lang][key]; ok && msg != "" {
		return msg
	}
	if msg, ok := catalogs["en"][key]; ok && msg != "" {
		return msg
	}
	return string(key)
}

// Available lists supported languages.
func Available() []Lang {
	return []Lang{
		{Code: "es", Name: "Español"},
		{Code: "en", Name: "English"},
		{Code: "fr", Name: "Français"},
		{Code: "de", Name: "Deutsch"},
		{Code: "it", Name: "Italiano"},
		{Code: "pt", Name: "Português"},
	}
}

// IsSupported reports whether lang is registered.
func IsSupported(lang string) bool {
	_, ok := catalogs[lang]
	return ok
}
