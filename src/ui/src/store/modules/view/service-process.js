const state = {
  localProcessTemplate: []
}

const getters = {
  localProcessTemplate: state => state.localProcessTemplate,
  // eslint-disable-next-line max-len
  hasProcessName: state => process => state.localProcessTemplate.find(template => template.bk_func_name.value === process.bk_func_name.value)
}

const actions = {}

const mutations = {
  setLocalProcessTemplate(state, processes) {
    state.localProcessTemplate = processes
  },
  addLocalProcessTemplate(state, process) {
    state.localProcessTemplate.push(process)
  },
  updateLocalProcessTemplate(state, { process, index }) {
    state.localProcessTemplate.splice(index, 1, process)
  },
  deleteLocalProcessTemplate(state, { index }) {
    state.localProcessTemplate.splice(index, 1)
  },
  clearLocalProcessTemplate(state) {
    state.localProcessTemplate = []
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}
