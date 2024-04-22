from django.contrib import admin
from django.urls import path
from Medico import views
from django.contrib.auth import views as auth_views
from django.conf import settings
from django.conf.urls.static import static


urlpatterns = [
    path('',views.Home, name='Home'),
    path('signup/',views.Signup, name='signup'),
    path('login/',views.Login, name='login'),
    path('Medicine/', views.Product_List, name='products'),
    path('chat/', views.chat_view, name='chats'),
    path('profile/',views.Profile, name='persons'),
    path('cart/', views.view_cart, name='view_cart'),
    path('cart/add/', views.add_to_cart, name='add_to_cart'),
    path('remove_from_cart/<int:product_id>/', views.remove_from_cart, name='remove_from_cart'),
    path('clear_cart/', views.clear_cart, name='clear_cart'),
    path('about-us/', views.About, name='aboutus'),
    path('blood/', views.Blood, name='blood'),
    path('blood/request/', views.RequestBlood, name='Requestblood'),
    path('blood/request/submit', views.RequestBlood, name='RequestbloodSubmit'),
    path('blood/donate/', views.DonateBlood, name='Donateblood'),
    path('blood/donate/submit', views.DonateBlood, name='DonatebloodSubmit'),
    path('logout/', auth_views.LogoutView.as_view(), name='logout'),
    path('profile/edit',views.edit_profile,name='editProfile'),
    path('search/', views.product_search, name='product_search'),
    path('admin/', admin.site.urls),
]

if settings.DEBUG:
    urlpatterns += static(settings.MEDIA_URL, document_root=settings.MEDIA_ROOT)
