import {useRoute, useRouter} from 'vue-router';
import {computed} from 'vue';

export const useAdminRoutes = () => {
    const route = useRoute();
    const router = useRouter();

    const entity = computed(() => route.query.entity ?? null);
    const action = computed(() => route.query.action ?? null);

    const setEntity = (newEntity) => {
        setAction(action.value === null ? null : "").then(
            () => router.replace({
                query: {
                    ...route.query ?? {},
                    entity: newEntity
                }
            })
        );
    };

    const setAction = (newAction) => {
        return setSelector(null).then(
            () => router.replace({
                query: {
                    ...route.query ?? {},
                    action: newAction
                }
            })
        );
    };

    // Currently this just lives here to make async promises happy
    const setSelector = (newSelector) => {
        return router.replace({
            query: {
                ...route.query ?? {},
                selector: newSelector
            }
        })
    };

    return {entity, action, setEntity, setAction};
};