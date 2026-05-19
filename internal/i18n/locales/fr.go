package locales

import "github.com/gsusgit/tuido/internal/i18n"

func init() {
	i18n.Register("fr", map[i18n.Key]string{
		i18n.KeyNewTask: "Nouveau", i18n.KeyFilter: "Filtrer", i18n.KeyFilterActive: "Filtre",
		i18n.KeyEdit: "Modifier", i18n.KeyDelete: "Supprimer", i18n.KeyComplete: "Terminer",
		i18n.KeyNavigate: "Naviguer", i18n.KeyTheme: "Thème", i18n.KeyExit: "Quitter",
		i18n.KeySearch: "Rechercher", i18n.KeyEmptyList: "Aucune tâche. Appuyez sur 'n'.",
		i18n.KeyEmptyFiltered: "Aucune tâche correspondante.",
		i18n.KeyTitleRequired: "Le titre ne peut pas être vide",
		i18n.KeyCompletedOf: "Terminées : %d sur %d", i18n.KeyScrollMore: "Voir plus (défiler ou %s)",
		i18n.KeyTitle: "Titre", i18n.KeyCategory: "Catégorie", i18n.KeyPriority: "Priorité",
		i18n.KeyStatus: "État", i18n.KeyPending: "En attente", i18n.KeyStatusCompleted: "Terminées",
		i18n.KeyFilterTitle: "Filtrer les tâches", i18n.KeyAll: "Toutes",
		i18n.KeyTabNext: "Suivant", i18n.KeyEnterSave: "Enregistrer", i18n.KeyEnterApply: "Appliquer",
		i18n.KeyEscCancel: "Annuler", i18n.KeyBack: "Retour", i18n.KeyReset: "Réinitialiser",
		i18n.KeyTerminalSmall:       "Terminal trop petit — redimensionnez",
		i18n.KeyWindowTooSmall:      "Fenêtre trop petite",
		i18n.KeyMinimumSizeRequired: "Taille minimale : %d×%d",
		i18n.KeyDeletePending: "Supprimer ? (Entrée confirmer · Esc annuler)",
		i18n.KeyTaskCreated: "Tâche créée", i18n.KeyTaskSaved: "Tâche enregistrée",
		i18n.KeyTaskDeleted: "Tâche supprimée", i18n.KeySaveError: "Échec : %s",
		i18n.KeyLangChanged: "Langue : %s", i18n.KeyResetPrompt: "Supprimer %s ? (o/N) : ",
		i18n.KeyResetDone: "Données supprimées.", i18n.KeyResetCancelled: "Annulé.",
		i18n.KeyResetInvalid: "Langue inconnue : es, en, fr, de, it, pt",
		i18n.KeyPersonal: "Personnel", i18n.KeyTrabajo: "Travail",
		i18n.KeyPriorityHigh: "Haute", i18n.KeyPriorityMedium: "Moyenne", i18n.KeyPriorityLow: "Basse",
		i18n.KeySelectLang: "Choisir la langue :", i18n.KeyLangNav: "↑↓ · Entrée · Esc",
		i18n.KeyControls: "Contrôles", i18n.KeyHelpMenu: "Voir les contrôles",
		i18n.KeySort: "Tri", i18n.KeySortField: "Trier par", i18n.KeySortDirection: "Direction",
		i18n.KeySortDefault: "Défaut", i18n.KeySortAsc: "Croissant", i18n.KeySortDesc: "Décroissant",
	})
}
