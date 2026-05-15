import { colors, typography } from "../../theme/design-system";

export default function Header() {
    return (
        <header style={{
            height: '65px',
            backgroundColor: colors.primary,
            display: 'flex',
            alignItems: 'center',
            justifyContent: 'center',
            padding: '0 24px',
        }}>
            <span style={{
                fontSize: typography.fontSize.xxl,
                color: colors.primaryLight
            }}>Industrix Todo App</span>
        </header>
    )
}