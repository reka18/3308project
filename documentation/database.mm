<map version="freeplane 1.6.0">
<!--To view this file, download free mind mapping software Freeplane from http://freeplane.sourceforge.net -->
<node TEXT="SocialMediaSite Database Table Map" FOLDED="false" ID="ID_654665779" CREATED="1569552254878" MODIFIED="1569554232446" STYLE="oval" VGAP_QUANTITY="0.0 pt">
<font SIZE="18"/>
<hook NAME="MapStyle">
    <properties edgeColorConfiguration="#808080ff,#ff0000ff,#0000ffff,#00ff00ff,#ff00ffff,#00ffffff,#7c0000ff,#00007cff,#007c00ff,#7c007cff,#007c7cff,#7c7c00ff" fit_to_viewport="false"/>

<map_styles>
<stylenode LOCALIZED_TEXT="styles.root_node" STYLE="oval" UNIFORM_SHAPE="true" VGAP_QUANTITY="24.0 pt">
<font SIZE="24"/>
<stylenode LOCALIZED_TEXT="styles.predefined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="default" ICON_SIZE="12.0 pt" COLOR="#000000" STYLE="fork">
<font NAME="SansSerif" SIZE="10" BOLD="false" ITALIC="false"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.details"/>
<stylenode LOCALIZED_TEXT="defaultstyle.attributes">
<font SIZE="9"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.note" COLOR="#000000" BACKGROUND_COLOR="#ffffff" TEXT_ALIGN="LEFT"/>
<stylenode LOCALIZED_TEXT="defaultstyle.floating">
<edge STYLE="hide_edge"/>
<cloud COLOR="#f0f0f0" SHAPE="ROUND_RECT"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.user-defined" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="styles.topic" COLOR="#18898b" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subtopic" COLOR="#cc3300" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subsubtopic" COLOR="#669900">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.important">
<icon BUILTIN="yes"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.AutomaticLayout" POSITION="right" STYLE="bubble">
<stylenode LOCALIZED_TEXT="AutomaticLayout.level.root" COLOR="#000000" STYLE="oval" SHAPE_HORIZONTAL_MARGIN="10.0 pt" SHAPE_VERTICAL_MARGIN="10.0 pt">
<font SIZE="18"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,1" COLOR="#0033ff">
<font SIZE="16"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,2" COLOR="#00b439">
<font SIZE="14"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,3" COLOR="#990000">
<font SIZE="12"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,4" COLOR="#111111">
<font SIZE="10"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,5"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,6"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,7"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,8"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,9"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,10"/>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,11"/>
</stylenode>
</stylenode>
</map_styles>
</hook>
<hook NAME="AutomaticEdgeColor" COUNTER="51" RULE="ON_BRANCH_CREATION"/>
<node TEXT="comments" POSITION="right" ID="ID_1448319805" CREATED="1569552951503" MODIFIED="1569554206742" HGAP_QUANTITY="294.49999164044885 pt" VSHIFT_QUANTITY="-100.49999700486669 pt">
<edge COLOR="#ff00ff"/>
<hook NAME="FreeNode"/>
<node TEXT="*comment_id" ID="ID_478230980" CREATED="1569552964038" MODIFIED="1569552996221"/>
<node TEXT="post_id" ID="ID_1574068631" CREATED="1569552972263" MODIFIED="1569552980467"/>
<node TEXT="user_id" ID="ID_275996129" CREATED="1569552980860" MODIFIED="1569553035579"/>
<node TEXT="text" ID="ID_1312155199" CREATED="1569553036384" MODIFIED="1569553039911"/>
<node TEXT="upvote[user_id]" ID="ID_644576547" CREATED="1569553040265" MODIFIED="1569553921931"/>
<node TEXT="downvote[user_id]" ID="ID_602732400" CREATED="1569553043493" MODIFIED="1569553931572"/>
<node TEXT="hidden" ID="ID_893290143" CREATED="1569553047055" MODIFIED="1569553054758"/>
</node>
<node TEXT="posts" POSITION="right" ID="ID_881807578" CREATED="1569552780017" MODIFIED="1569554232445" HGAP_QUANTITY="382.9999890029434 pt" VSHIFT_QUANTITY="2.2499999329447746 pt">
<edge COLOR="#00ff00"/>
<hook NAME="FreeNode"/>
<node TEXT="*post_id" ID="ID_1140340873" CREATED="1569552867403" MODIFIED="1569552990077"/>
<node TEXT="user_id" ID="ID_1965311537" CREATED="1569553008351" MODIFIED="1569553014644"/>
<node TEXT="text" ID="ID_409064162" CREATED="1569552895930" MODIFIED="1569552910210"/>
<node TEXT="upvotes[user_id]" ID="ID_840954535" CREATED="1569552910968" MODIFIED="1569553894412"/>
<node TEXT="downvotes[user_id]" ID="ID_796458531" CREATED="1569552919141" MODIFIED="1569553902645"/>
</node>
<node TEXT="users" POSITION="right" ID="ID_1480922463" CREATED="1569552753552" MODIFIED="1569554193310" HGAP_QUANTITY="321.49999083578615 pt" VSHIFT_QUANTITY="127.499996200204 pt">
<edge COLOR="#00ffff"/>
<hook NAME="FreeNode"/>
<node TEXT="*user_id" ID="ID_31290674" CREATED="1569552785708" MODIFIED="1569553000984"/>
<node TEXT="first_name" ID="ID_953326818" CREATED="1569552839142" MODIFIED="1569553261758"/>
<node TEXT="last_name" ID="ID_1802067914" CREATED="1569552843888" MODIFIED="1569553264562"/>
<node TEXT="email" ID="ID_1689316793" CREATED="1569552846614" MODIFIED="1569552849334"/>
<node TEXT="gender" ID="ID_1124380053" CREATED="1569552849732" MODIFIED="1569552852490"/>
<node TEXT="password" ID="ID_751831146" CREATED="1569553219219" MODIFIED="1569553222928"/>
<node TEXT="follows[user_id]" ID="ID_287525569" CREATED="1569553778975" MODIFIED="1569553837540"/>
</node>
<node TEXT="avatars" POSITION="left" ID="ID_314295067" CREATED="1569553245041" MODIFIED="1569554153516" HGAP_QUANTITY="53.74999881535771 pt" VSHIFT_QUANTITY="44.99999865889552 pt">
<edge COLOR="#007c00"/>
<node TEXT="*avatar_id" ID="ID_1036847832" CREATED="1569553250001" MODIFIED="1569553416731"/>
<node TEXT="user_id" ID="ID_1118379761" CREATED="1569553714119" MODIFIED="1569553719202"/>
<node TEXT="binary_data" ID="ID_148693959" CREATED="1569553272250" MODIFIED="1569553733272"/>
</node>
</node>
</map>
