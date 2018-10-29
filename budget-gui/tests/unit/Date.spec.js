import { shallowMount, createLocalVue } from '@vue/test-utils'
import DateComponent from '@/components/Date.vue'
import Vuex from 'vuex'
import ShortKey from 'vue-shortkey'
import ElementUI from 'element-ui';

const localVue = createLocalVue()
localVue.use(Vuex)
localVue.use(ShortKey)
localVue.use(ElementUI)

let store = new Vuex.Store({
  state: {
    date: {
      day: 1,
      month: 1,
      year: 2000,
    },
  },
  mutations: {
    setDay(state, day) {
      state.date.day = day;
    },
    setMonth(state, month) {
      state.date.month = month;
    },
    setYear(state, year) {
      state.date.year = year;
    }
  },
})

describe('Date.vue', () => {
  it('Increment to february 29 2000', () => {
    const wrapper = shallowMount(DateComponent, { store, localVue });
    wrapper.vm.day = 28;
    wrapper.vm.month = 2;
    wrapper.vm.year = 2000;
    wrapper.vm.increment();
    expect(wrapper.vm.day).toBe(29);
  })

  it('Decrement to february 29 2000', () => {
    const wrapper = shallowMount(DateComponent, { store, localVue });
    wrapper.vm.day = 1;
    wrapper.vm.month = 3;
    wrapper.vm.year = 2000;
    wrapper.vm.decrement();
    expect(wrapper.vm.day).toBe(29);
  })

  it('Increment from february 28 to march 1 2001', () => {
    const wrapper = shallowMount(DateComponent, { store, localVue });
    wrapper.vm.day = 28;
    wrapper.vm.month = 2;
    wrapper.vm.year = 2001;
    wrapper.vm.increment();
    expect(wrapper.vm.day).toBe(1);
    expect(wrapper.vm.month).toBe(3);
  })

  it('Decrement from march 1 to february 28 2001', () => {
    const wrapper = shallowMount(DateComponent, { store, localVue });
    wrapper.vm.day = 1;
    wrapper.vm.month = 3;
    wrapper.vm.year = 2001;
    wrapper.vm.decrement();
    expect(wrapper.vm.day).toBe(28);
    expect(wrapper.vm.month).toBe(2);
  })

  it('Increment from december 31 2000 to january 1 2001', () => {
    const wrapper = shallowMount(DateComponent, { store, localVue });
    wrapper.vm.day = 31;
    wrapper.vm.month = 12;
    wrapper.vm.year = 2000;
    wrapper.vm.increment();
    expect(wrapper.vm.day).toBe(1);
    expect(wrapper.vm.month).toBe(1);
    expect(wrapper.vm.year).toBe(2001);
  })

  it('Decrement from january 1 2001 to december 31 2000', () => {
    const wrapper = shallowMount(DateComponent, { store, localVue });
    wrapper.vm.day = 1;
    wrapper.vm.month = 1;
    wrapper.vm.year = 2001;
    wrapper.vm.decrement();
    expect(wrapper.vm.day).toBe(31);
    expect(wrapper.vm.month).toBe(12);
    expect(wrapper.vm.year).toBe(2000);
  })
})
