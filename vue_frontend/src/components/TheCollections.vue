<template>
    <div class="collections">

        <div class="collections__item" v-for="collection in getCollections" :key="collection.id"
            :class="{ 'collections__item--selected': collection.id == getActiveCollection }"
            @click="selectCollection(Number(collection.id))">


            <!-- Name or edit field -->
            <div class="collections__main">
                <template v-if="editingId === Number(collection.id)">
                    <input v-model="editingName" class="collections__input" type="text"
                        @keyup.enter="saveEdit(collection)" @keyup.esc="cancelEditDelete" @blur="saveEdit(collection)"
                        autofocus />
                </template>

                <template v-else>
                    <span class="collections__name">{{ collection.name }}</span>
                </template>
            </div>

            <!-- Action buttons -->
            <div class="collections__actions" v-if="collection.id == getActiveCollection && collection.id != -1">
                <template v-if="editingId === Number(collection.id) || deletingId === Number(collection.id)">
                    <svg class="collections__btn collections__btn--ok" @click.stop="editOrDelete(collection)">
                        <use xlink:href="@/ui/svg/sprite.svg#icon-very-good"></use>
                    </svg>
                    <svg class="collections__btn collections__btn--cancel" @click.stop="cancelEditDelete">
                        <use xlink:href="@/ui/svg/sprite.svg#icon-close-bold"></use>
                    </svg>
                </template>

                <template v-else>
                    <svg class="collections__btn collections__btn--edit" @click.stop="startEdit(collection)">
                        <use xlink:href="@/ui/svg/sprite.svg#icon-edit"></use>
                    </svg>
                    <svg class="collections__btn collections__btn--delete" @click.stop="startDelete(collection)">
                        <use xlink:href="@/ui/svg/sprite.svg#icon-delete"></use>
                    </svg>
                </template>
            </div>
        </div>

    </div>
</template>


<!-- --------------------------------------------------------------- -->

<script setup lang="ts">
import { ref, watch } from "vue";
import { useChessLogStore } from "@/stores/chessLogStore";
import { storeToRefs } from "pinia";
import { useRoute, useRouter } from "vue-router";
import { useAppStore } from "@/stores/appStore";
import axios from "@/axios";

const router = useRouter();
const route = useRoute();

const chessLogStore = useChessLogStore();
const { getCollections, getActiveCollection } = storeToRefs(chessLogStore);

const appStore = useAppStore();

const editingId = ref<number | null>(null);
const deletingId = ref<number | null>(null);
const editingName = ref("");



// - Wachers -----------------------------------------------------------

watch(
    getActiveCollection,
    () => {
        cancelEditDelete()
    }
);

function selectCollection(id: number) {
    chessLogStore.setActiveCollection(id);
}



function startEdit(collection: any) {
    deletingId.value = null;
    editingId.value = Number(collection.id);
    editingName.value = collection.name;
}

async function saveEdit(collection: any) {
    if (editingId.value !== Number(collection.id)) return;

    try {
        
        const url = `${appStore.getAppUrl}/api/collections/${collection.id}`;

        const newName = editingName.value.trim();
        if (!newName) return cancelEditDelete();

        const r = await axios.put(url, {name: newName});

        if (r.statusText == "OK") {
            collection.name = newName;    
            editingId.value = null;
            editingName.value = "";
        }

    } catch (err) {
        console.error("! TheCollections saveEdit !\n", err)
    }
}


function startDelete(collection: any) {
    editingId.value = null;
    deletingId.value = Number(collection.id);
    editingName.value = "";
}

async function deleteCollection(collection: any) {

    try {

        const url = `${appStore.getAppUrl}/api/collections/${collection.id}`;

        const r = await axios.delete(url);

        if (r.statusText == "OK") {


            const index = getCollections.value.findIndex(c => c.id == collection.id);
            if (index !== -1) getCollections.value.splice(index, 1);
    
            // If selected collection is deleted, reset active collection
            if (getActiveCollection.value === collection.id) {
                chessLogStore.setActiveCollection(0);
            }
    
            deletingId.value = null;

        }
  

    } catch (err) {
        console.error("! deleteCollection ! \n", err)
    }

}

function editOrDelete(collection: any) {
    if (deletingId.value) return deleteCollection(collection);
    if (editingId.value) return saveEdit(collection);
}

function cancelEditDelete() {
    editingId.value = null;
    deletingId.value = null;
    editingName.value = "";
}
</script>


<!-- --------------------------------------------------------------- -->

<style lang="scss" scoped>
.collections {
    font-family: var(--font-primary);
    color: var(--color-text-1);
    font-size: .9rem;


    &__item {
        padding: .5rem;
        cursor: pointer;
        display: grid;
        grid-template-columns: 1fr 3rem;
        align-items: center;
        justify-content: space-between;
        height: 3rem;
        gap: 5px;
        transition: background-color .2s;
        



        &:hover {
            background-color: var(--color-slate-50);

            .collections__actions {
                opacity: 1;
                pointer-events: auto;
            }
        }

        &--selected {
            background-color: var(--color-slate-100);
            font-weight: 600;
        }
    }

    &__color {
        border-left: 5px solid;
        border-radius: 10px;
        height: 100%;
        margin-right: .4rem;
    }

    &__main {
        flex: 1;
        max-width: 10rem;
    }

    &__name {
        text-wrap: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        padding: 2px 2px 2px 5px;
    }

    &__input {
        width: 100%;
        font: inherit;
        padding: 2px 2px 2px 4px;
        border-radius: 4px;
        border: 1px solid var(--color-slate-300);
        background-color: var(--color-slate-50);
        color: var(--color-slate-700);

        &:focus {
            outline: none;
            border-color: var(--color-slate-500);
        }
    }

    &__actions {
        display: flex;
        gap: .3rem;
        opacity: 0;
        pointer-events: none;
        transition: .2s;
        


        z-index: 100;
        height: 1.2rem;
        display: flex;
        align-items: center;
        // padding: 0px 5px;
        gap: 5px;
    }

    &__btn {        
        cursor: pointer;
        width: 1.4rem;
        height: 1.4rem;
        padding: 3px;
        background-color:  var(--color-slate-100);
      
        color: var(--color-slate-700);
        fill: currentColor;

        border-radius: 5px;
        border: 1px solid currentColor;

        &:hover {
            color: var(--color-rose-500);
        }
    }
}
</style>
