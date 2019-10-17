package handler

import (
	"encoding/json"
	"net/http"

	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

func New(reconciler reconcile.Reconciler) *Handler {
	return &Handler{reconciler: reconciler}
}

type Handler struct {
	reconciler reconcile.Reconciler
}

func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Add("Accept", http.MethodPost)
		http.Error(w, "Method not allowed", 405)
		return
	}

	nn := types.NamespacedName{}
	err := json.NewDecoder(r.Body).Decode(&nn)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	request := reconcile.Request{NamespacedName: nn}
	result, err := h.reconciler.Reconcile(request)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	json.NewEncoder(w).Encode(&result)
}
